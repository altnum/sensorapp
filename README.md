# Sensor App
Sensor App is composed of a CLI tool, REST server and two DBs. Its purpose is to give
information for different system stats like CPU usage and temperature and Memory utilization.
The user can choose between different sensor groups (configured with the model.yaml
file) which provide various sensors with specified units. The output from the program is
being shown in the terminal, in the output.csv file (if given path) and in a database (if
given URL request to the server). Except from the value and the provided "deviceid" and "sensorid"
in the measurement, each output provides a timestamp so as to check when the
measurement took place exactly. All the outputs are written simultaneously.
The project uses Cobra for the CLI interactions, Mux for the request router, Mockery for mocking, 
Ginkgo and Gomega for the unit tests, PostgreSQL for
keeping the different devices and sensors and InfluxDB for the measurements and Docker for
building and running the database images.
## How to build and run the project
After setting the Go environment and pulling the repository:  
1) build your binary files- run `$ go build` in '~/cli' and in '\~/server' directories  
2) build your database Docker images:  
run `$ docker build -t influxdb .` in '~/influxdb' and `$ docker build -t postgresql .` in '\~/postgresql'  
3) run your databases on the respective ports  
`$ docker run -p 5432:5432 postgresql` and `$ docker run -p 8086:8086 influxdb`  
4) run your server- `$ ./sensor` in the '~/server' dir  
5) To start your CLI app with the desired commands run `$ ./sensor` in the '~/cli' dir.
## How to use the given commands  
<li><b>--unit</b> is used to specify your unit preference- C/F (Celsius/Fahrenheit). It is used just for the CPU_TEMP sensor group</li>
<li><b>--format</b> is used to specify the preferred output type- JSON/YAML</li>
<li><b>--total_duration</b> is used to specify the preferred duration time of the program in seconds</li>
<li><b>--delta_duration</b> is used to specify the time passing between two sensor measurements in seconds</li>
<li><b>--sensor_group</b> is used to specify what kind of sensor you want to use</li>
<li><b>--output_file</b> is used to specify the destination to the desired .csv output file</li>
<li><b>--web_hook_url</b> is used to specify the server URL for persisting the measurements in the InfluxDB database</li>