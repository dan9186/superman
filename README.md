# Superman
[![Build](https://github.com/dan9186/superman/workflows/Build/badge.svg?branch=master)](https://github.com/dan9186/superman/actions?query=workflow%3ABuild)
[![Go Reportcard](https://goreportcard.com/badge/github.com/dan9186/superman)](https://goreportcard.com/report/github.com/dan9186/superman)

Superman takes in login events associated with a unix timestamp and an IP address. The provided information is used to reference a geolocation and provide back details of suspicous logins based on login activity over time.

## The Problem

Given mulitiple login events from a single user, a user wants to know when their login events are not posible given normal human means. That is to say that if two login events are geographically separated by a distance and time that would require greater than 500 mph of travel, it is a high probability the same person didn't perform both login events and their account has been compromised.

## Building
All steps below are run the same way they are in the build. The build may chain the steps slightly differently for the benefit of reporting, but otherwise does the same thing.

### Running

The service will be stood up in a local docker environment and bound to `localhost:4567` for allowing any manual poking and prodding of the service.

```
make run
```

It is worth noting that multiple calls to `make run` will result in the service being rebuilt and freshly started. So the development pattern can be `code => make run test => code => make run test => ...`

### Test

The test target executes both unit tests and functional tests. If the local environment of the service is not running, the functional tests will fail.

```
make test
```

### Clean

To destroy the local environment and remove any other generated files, you can clean the project and start fresh.

```
make clean
```

### Help

The help target will provide additional targets not highlighted above along with explanations of what each does.

```
make help
```


## Resources Used

* [Gomicro Cucumber Docker Image](https://github.com/gomicro/docker-cucumber) - used to execute BDD style tests written in cucumber
* [Gomicro Service](https://github.com/gomicro/service) - used to bootstrap service (dockerfile, docker-compose, minimal running go service)
* [UUID](https://github.com/google/uuid) - used to facilitate uuid parsing
* [Goblin](https://github.com/franela/goblin) - used for organizing unit tests and better report output
* [Gomega](https://github.com/onsi/gomega) - used for a convenient set of assertion matchers in tests
* [SQLMock](https://github.com/DATA-DOG/go-sqlmock) - used for mocking the database and enable unit testing of database calls
* [GeoIP2]("github.com/oschwald/geoip2-golang") - used for reading the geoip2 data and parsing results

## Notes

### SQLite Requirement

Normally I would not use SQLite even for local development, as a significantly closer to deployed code setup can be accomplished with a postgres container.

Specific Items of Note:

* SQLite requires `CGO` to be enabled and requires either installing build tools or using the non `alpine` Golang image
* The build on a fresh build instance takes longer due to the size difference of the `golang-alpine` vs `golang` docker images (370MB vs 809MB)
* The resulting docker image is also larger, because we cannot run a scratch container (25MB vs 99MB)
* Bootstrapping of the database gets built into the binary rather than being able to build it into the database image or use some migration tool that would apply to prod
