# About
gofrankie is a program written in Golang to demonstrate building REST service in Gin framework.

# Key Technical Points
- Tested with Golang 1.14.1
- Use Gin as the REST framework
- 

# Assumptions
- Poor or even No CSV content validation
- CSV file format is correct
- No timezone consideration, just assume local timezone
- No input validation (e.g. Threshold argument)
- No consideration of currency, assuming same currency for all transactions
- No need to suppress alerts for the same credit card hash.
    * In other words, if a transaction has triggered an alert, a following transaction from the same hash will likely also trigger another alert (as long as sliding window total is over threshold)
- Done a feature which might not be necessary but personally feel would of of interest
    * When an alert is trigger, report the amount of the window, and number of transactions in the window.

# Design Considerations and Decisions
- Java Streaming
    * map/filter/collect
    * Easier to read
- Algorithm
    * Basically few steps as below
        * Group by hash, result of this is that each hash yields a list of transactions belonging to same hash
        * Use a moving/sliding window from first transaction to last to analyse transactions of same hash. Collect possible alerts while moving the window.
        * Collect all alerts

# How to Build
Assuming Go/Docker is already installed
- Non Docker Approach
    * Clone `git clone git@github.com:alfredxiao/gofrankie.git`
    * `cd gofrankie`
    * Build `go build`, which should create an executable `gofrankie`
- Docker Approach
    * run `git clone` and `cd` Same as above
    * run `./docker_build.sh`, which should create an docker image `gofrankie:latest`

# How to Run
- Non Docker
    * run `./gofrankie`, it is hardcoded to listen on port 8080 
- Docker
    * `./docker_run.sh`

# Test the API
Once `gofrankie` is started, we can test its `isgood` API.

```
curl -X POST 'http://localhost:8080/isgood' --header 'Content-Type: application/json' \
--data-raw '[
  {
    "activityType": "SIGNUP",
    "checkSessionKey": "key100",
    "checkType": "COMBO",
    "activityData": [
      {
        "kvpKey": "name",
        "kvpValue": "alfred",
        "kvpType": "general.string"
      }
    ]
  }
]'
```

# Room for Improvement
- Error Codes
    * Define finer grained error codes for various specific errors
- 100% coverage
    * Make coverage 100 percent, covering main as well
- Not hardcode the port number
    * Accept from command argument or environment variable?