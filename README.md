# Solution 
Solution to take home code test.

## Story
There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

## Goal
To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

## Api
Exposed API is a restful [Open API](./api/open-api.yaml) and it provides two paths.

### Calculate
Calculates source and final destination of a person's flight.
```http
POST http://localhost:8080/calculate HTTP/1.1
content-type: application/json

[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]  

```
### Itinerary
Returns ordered list flights taken by the person. 
```http
POST http://localhost:8080/itinerary HTTP/1.1
content-type: application/json

[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]  

```

## Test Build Run

* Pre-requisites
  - Go SDK
  - Makefile
  - Docker (Optional)
  - VS Code (Optional)
  - Curl (Either VS Code)

### Unit Tests
```bash
make test
```

There are few ways this service can be run and tested. In all cases service will start and listen on port 8080

### Run Local
```bash
make serve # this will run go web service locally
```

### Run Docker Container
* Pre-requisites
  - Docker Setup
  - Port 8080 is free
  
Build image and run a container locally.
```bash
make pack # builds a docker image name - `flyer`
make run # runs flyer as a docker container
```

### Run Docker Compose
Running using docker compose gives swagger-ui experience to view open api spec and test using browser.
* Pre-requisites
  - Docker Setup
  - Ports 8080, 7080 are free

This will start flyer and swagger-ui containers. After successful startup
```bash
make compose # on successful startup should open swagger-ui[http://localhost:7080] browser
```
> Note: If port 7080 is unavailable, please edit following files with available host port

> - [docker-compose](./docker-compose.yaml) # locate swagger-ui and change host port
> - [Makefile](./Makefile) # locate compose target and change port

### HTTP Tests

### VS Code Extention: REST Client
* Pre-requisites
  - VS Code
  - Extention: REST Client
Open file [example http](example.http) and click on 'Send Request' on any of the samples.

### Curl
```bash
curl -v -X POST http://localhost:8080/calculate \
   -H 'Content-Type: application/json' \
   -d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'
```

## Cleanup
```bash
make clean # removes any binaries or docker containers built
```