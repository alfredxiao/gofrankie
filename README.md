# About
`gofrankie` is a program written in Golang to demonstrate building REST service in **Gin** framework.

# Key Technical Points
- Tested with **Golang 1.14.1**
- Uses **Gin** as the REST framework: https://gin-gonic.com/
- Does NOT use tool to generate models from swagger doc

# Assumptions
- We do not have full list of vendor specific activity type for field `ActivityType`, rather we just treat anything starting with `_` as vendor specific.
- API client related errors are returned as 400 Bad Request error
- Assume the default handling from Gin for error cases - e.g. route not found 404 error, internal server error 500 - is acceptable.
- No Authentication required
- No TLS/SSL required

# Design Considerations and Decisions
- Try to separate logic/code into different pacakges while leaving each package not too large
- Utilises Gin framework's validation as well as custom extension to validate data
- Chooses Gin over chi, for its relative simpler/quicker setup, e.g. JSON
    * Reference: https://brunoscheufler.com/blog/2019-04-26-choosing-the-right-go-web-framework
- When testing, try testing at unit level rather than system integration level

# How to Build
Assuming Go/Docker is already installed
- Non Docker Approach
    * Clone `git clone git@github.com:alfredxiao/gofrankie.git`
    * `cd gofrankie`
    * Build `go build`, which should create an executable `gofrankie`
- Docker Approach
    * run `git clone` and `cd` Same as above
    * run `./docker_build.sh`, which should create an docker image `gofrankie:latest`

# Test
- Run `go test ./...`
- Or, Run `./test.sh` which will generate a coverage report and open it in a browser

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

# Notes
- `server.go` has been made this way with the intent to be able to shutdown down it gracefully from both when running in test and from when interrupt signal, e.g. Ctrl-C. Otherwise, it could be made much simpler, with even 20 lines fewer.
- Only `server_test.go` requires running real server, other tests are unit level and require no server running
- `vendor` folder is not checked, if needed, one can just run `go mod vendor` to recreate them.

# Room for Improvement
- Error Codes
    * Define finer grained error codes for various specific errors
- Return all validation issues in one go rather than return just the first issue encountered
    * E.g. If request data has two fields not conforming to specification, tell consumers both issues
- 100% test coverage
    * Make coverage 100 percent, covering main as well
- Not hardcode the port number
    * Accept from command argument or environment variable?
- Make docker image more secure
    * e.g. not run as root user