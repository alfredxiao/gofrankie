# About
gofrankie is a program written in Golang to demonstrate building REST service in Gin framework.

# Key Technical Points
- Tested with Golang 1.14.1
- Use Gin as the REST framework
- Does NOT use tool to generate models from swagger doc

# Assumptions
- ??

# Design Considerations and Decisions
- ??

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