# Sensor App
Sensor App is composed of a CLI tool, REST server, two DBs and configurable dashboards
for interface, provided by Grafana. Its purpose is to give
information for different system stats like CPU usage and temperature and Memory utilization.
The user can choose between different sensor groups (configured with the model.yaml
file) which provide various sensors with specified units. The output from the program is
being shown in the terminal, in the output.csv file (if given path) and in a database (if
given URL request to the server). Except from the value and the provided "deviceid" and "sensorid"
in the measurement, each output provides a timestamp so as to check when the
measurement took place exactly. All the outputs are written simultaneously.
The project uses Cobra for the CLI interactions, Mux for the request router, Mockery for mocking, 
Ginkgo and Gomega for the unit tests, PostgreSQL for
keeping the different devices and sensors, InfluxDB for the measurements and Docker for
building and running the database images. Docker is used for running the Grafana tool as well,
used for monitoring the different sensors output data and sending alert notifications if some
sensors values are out of the predefined boundaries.
For testing purposes, MailDev is added to the "docker-compose" configuration file- it is used for 
receiving the email alert notifications.
## How to build and run the project
After setting the Go environment and pulling the repository:  
1) build your binary files- run `$ go build` in '~/cli' and in '\~/server' directories  
2) build your PostgreSQL Docker image and run the containers needed:  
run `$ docker build -t postgresql .` in '\~/postgresql' `$ docker build -t server .` in '\~/server' and then `$ docker-compose up -d` in '~/grafana'  
3) run your server- `$ ./sensor` in the '~/server' dir  
4) To start your CLI app with the desired commands run `$ ./sensor` in the '~/cli' dir.
## How to use the given commands  
<li><b>--unit</b> is used to specify your unit preference- C/F (Celsius/Fahrenheit). It is used just for the CPU_TEMP sensor group</li>
<li><b>--format</b> is used to specify the preferred output type- JSON/YAML</li>
<li><b>--total_duration</b> is used to specify the preferred duration time of the program in seconds</li>
<li><b>--delta_duration</b> is used to specify the time passing between two sensor measurements in seconds</li>
<li><b>--sensor_group</b> is used to specify what kind of sensor you want to use</li>
<li><b>--output_file</b> is used to specify the destination to the desired .csv output file</li>
<li><b>--web_hook_url</b> is used to specify the server URL for persisting the measurements in the InfluxDB database</li>
