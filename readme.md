# PayerPoints WebService
This webservice was developed using the Go Gin framework. In order to run the webservice locally Go will have to be
installed. The official instructions can be found here: [Go Install Instructions](https://go.dev/doc/install).
Instructions will vary by system.

### Run Project
Once Go is installed and 'go version' gives correct response in order to run and test:
* open a command line terminal and navigate to root directory of this project.
* start go app using this command: `go run main.go`
* you will see a listener for requests locally

### Test
Once project is running there are 3 endpoints which can be used. These can be tested through the unit tests or a network request
my favorites are usually Postman or Curl (endpoints listed below).

* Unit Tests:
  * to test by running unit tests run `go test` from the project route, all the apis are tested and responses should be valid 201/202 responses

### Endpoints
If ran with default settings the base url should be `http://localhost:8080` followed by one of the endpoints below

#### Post
* `/addTransaction`
    * expecting: 
      * `{ "payer": "DANNON", "points": 1000, "timestamp": "2020-11-02T14:00:00Z" }`
    * response body: 
      * `{ "payer": "DANNON", "points": 1000, "timestamp": "2020-11-02T14:00:00Z" }`
#### Get
* `/getBalances`
    * response body: 
      * `{"DANNON": 1000,"UNILEVER": 0,"MILLER COORS": 5300}`
#### Put
* `/spendPoints`
    * expect: 
      * `{ "points": 5000 }`
    * response body:
      * `[{ "payer": "DANNON", "points": -100 },{ "payer": "UNILEVER", "points": -200 },{ "payer": "MILLER COORS", "points": -4,700 }]`



