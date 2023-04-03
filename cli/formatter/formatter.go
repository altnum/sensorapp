package formatter

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/spf13/cobra"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
	"gopkg.in/yaml.v2"
)

var getFlagValue = getFlagValueString
var jsonMarshal = json.Marshal
var yamlMarshal = yaml.Marshal

type IFormatter interface {
	FormatOutput(*cobra.Command, systeminfo.Measurement) (outputStr string, err error)
	FormatInfoOutput(string, systeminfo.Measurement) (string, error)
}

type Formatter struct{}

func getFlagValueString(cmd *cobra.Command, flag string) (string, error) {
	flag, err := cmd.Flags().GetString("format")
	if err != nil {
		return "", err
	}

	return flag, nil
}

//Format the output in either JSON or YAML format.
func (f *Formatter) FormatOutput(cmd *cobra.Command, measurement systeminfo.Measurement) (outputStr string, err error) {
	formatF, err := getFlagValue(cmd, "format")
	if err != nil {
		return "", err
	}

	if formatF != "" {
		finalStr := ""
		outputStr, err = f.FormatInfoOutput(formatF, measurement)
		if err != nil {
			return "", err
		}
		finalStr += outputStr + "\n"

		return finalStr, nil
	}

	return "", errors.New("invalid input for 'format' command")
}

func (f *Formatter) FormatInfoOutput(format string, measurement systeminfo.Measurement) (string, error) {
	format = strings.ToLower(format)

	if format == "json" {
		toPrint, err := jsonMarshal(measurement)
		if err != nil {
			return "", err
		}
		return string(toPrint), nil
	}
	if format == "yaml" {
		toPrint, err := yamlMarshal(measurement)
		if err != nil {
			return "", err
		}
		return string(toPrint), nil
	}

	return "", errors.New("invalid input for 'format' command")
}
