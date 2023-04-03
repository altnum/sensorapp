/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	executiontimer "github.wdf.sap.corp/I554249/sensor/execution_timer"
	"github.wdf.sap.corp/I554249/sensor/formatter"
	"github.wdf.sap.corp/I554249/sensor/logger"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
	"github.wdf.sap.corp/I554249/sensor/system_info/sensors"
	sensorsfactory "github.wdf.sap.corp/I554249/sensor/system_info/sensors_factory"
	"github.wdf.sap.corp/I554249/sensor/writer"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile        string
	wg             sync.WaitGroup
	total_duration int64
	delta_duration int64
	format         formatter.Formatter
	getFlagInfo    func(string) (string, error)
	sensorFactory  sensorsfactory.ISensorFactory
	outWriter      writer.IWriter
	timer          executiontimer.ITimer
)

var errorLogger = logger.GetLogger().Error
var infoLogger = logger.GetLogger().Info
var warningLogger = logger.GetLogger().Warning

var rootCmd = &cobra.Command{
	Use:   "sensor",
	Short: "The app measures system sensor values of your local laptop.",
	Long: `Using Cobra for the CLI flags, we are capable of creating easily this simple app- "sensor".
The app is measuring the sensor values of your laptop, giving you the opportunity to change
some of the options by your own preferences.`,

	Run: func(cmd *cobra.Command, args []string) {
		runtime.GOMAXPROCS(5)
		getFlagInfo = cmd.Flags().GetString

		sensors := prepareInstances(cmd)

		go func() {
			err := executeProgram(cmd, sensors, &format)
			if err != nil {
				errorLogger.Fatalln(err)
			}
		}()

		select {
		case <-time.After(time.Duration(total_duration) * time.Second):
			infoLogger.Println("Measurements finished.")
			os.Exit(0)
		}
	},
}

func prepareInstances(cmd *cobra.Command) []sensors.SystemSensor {
	sensorFactory = &sensorsfactory.SensorFactory{}
	sensors := createSensors(cmd)
	outWriter = &writer.Writer{}

	timer = &executiontimer.Timer{}
	err := setUpTimer(cmd)
	if err != nil {
		errorLogger.Fatalln(err)
	}

	return sensors
}

func setUpTimer(cmd *cobra.Command) error {
	err := timer.SetUpTimer(cmd)
	if err != nil {
		return err
	}
	total_duration = timer.GetTotalDuration()
	delta_duration = timer.GetDeltaDuration()

	return nil
}

func createSensors(cmd *cobra.Command) []sensors.SystemSensor {
	var err error
	var systemSensors []sensors.SystemSensor

	systemSensors, err = sensorFactory.SensorFactory(cmd)
	if err != nil {
		errorLogger.Fatalln(err)
	}

	return systemSensors
}

func executeProgram(cmd *cobra.Command, systemSensor []sensors.SystemSensor,
	format formatter.IFormatter) error {

	for {
		measurements, err := getSystemInfo(cmd, systemSensor)
		if err != nil {
			return err
		}

		for _, measurement := range measurements {
			outputStr, err := buildOutput(cmd, measurement, format)
			if err != nil {
				return err
			}

			err = writeOutput(cmd, outputStr, measurement)
			if err != nil {
				return err
			}
		}

		//the execution of code stops, so as to satisfy the delta_duration periods
		select {
		case <-time.After(time.Duration(delta_duration) * time.Second):
			continue
		}
	}
}

//Take the measurements, according to cmd.Flags() values and returns the output in the proper format.
func getSystemInfo(cmd *cobra.Command, sensors []sensors.SystemSensor) ([]systeminfo.Measurement, error) {
	measurements := make([]systeminfo.Measurement, 0)

	unitF, err := getFlagInfo("unit")
	if err != nil {
		return measurements, err
	}

	if unitF != "" {
		warningLogger.Println("'unit' command value will be ignored for CPU_USAGE and MEMORY_USAGE")
	}

	for _, sensor := range sensors {
		measurement, err := sensor.StartMeasurement(unitF)
		if err != nil {
			return measurements, err
		}
		measurements = append(measurements, measurement)
	}

	return measurements, nil
}

func buildOutput(cmd *cobra.Command, measurement systeminfo.Measurement, format formatter.IFormatter) (string, error) {
	outputStr, err := buildOutputStr(cmd, measurement, format)
	if err != nil {
		return "", err
	}

	return outputStr, nil
}

func buildOutputStr(cmd *cobra.Command, measurement systeminfo.Measurement, format formatter.IFormatter) (string, error) {
	var outputStr string
	var err error
	outputStr, err = format.FormatOutput(cmd, measurement)
	if err != nil {
		return "", err
	}
	return outputStr, nil
}

func writeOutput(cmd *cobra.Command, outputStr string, measurement systeminfo.Measurement) error {
	outputFile, err := getFlagInfo("output_file")
	if err != nil {
		return err
	}

	serverURL, err := getFlagInfo("web_hook_url")
	if err != nil {
		return err
	}

	err = outWriter.StartWriting(serverURL, measurement, outputStr, outputFile)
	if err != nil {
		return err
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sensor.yaml)")
	rootCmd.PersistentFlags().String("unit", "", "Use --unit to specify your unit preference- C/F (Celsius/Fahrenheit)")
	rootCmd.PersistentFlags().String("format", "", "Use --format to specify the preferred output type- JSON/YAML")
	rootCmd.PersistentFlags().String("total_duration", "", "Use --total_duration to specify the preferred duration time of the program in seconds")
	rootCmd.PersistentFlags().String("delta_duration", "", "Use --delta_duration to specify the time passing between two sensor measurements in seconds")
	rootCmd.PersistentFlags().String("sensor_group", "", "Use --sensor_group to specify what kind of sensor you want to use")
	rootCmd.PersistentFlags().String("output_file", "", "Use --output_file to specify the destination to the desired .csv output file")
	rootCmd.PersistentFlags().String("web_hook_url", "", "Use --web_hook_url to specify the server URL for persisting the measurements in database")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sensor")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
