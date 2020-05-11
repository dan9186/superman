# Superman
[![Build](https://github.com/dan9186/superman/workflows/Build/badge.svg?branch=master)](https://github.com/dan9186/superman/actions?query=workflow%3ABuild)
[![Go Reportcard](https://goreportcard.com/badge/github.com/dan9186/superman)](https://goreportcard.com/report/github.com/dan9186/superman)

Superman takes in login events associated with a unix timestamp and an IP address. The provided information is used to reference a geolocation and provide back details of suspicious logins based on login activity over time.

## The Problem

Given multiple login events from a single user, a user wants to know when their login events are not possible given normal human means. That is to say that if two login events are geographically separated by a distance and time that would require greater than 500 mph of travel, it is a high probability the same person didn't perform both login events and their account has been compromised.

## Development
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

* [GeoIP2]("github.com/oschwald/geoip2-golang") - used for reading the geoip2 data and parsing results
* [Goblin](https://github.com/franela/goblin) - used for organizing unit tests and better report output
* [Gomega](https://github.com/onsi/gomega) - used for a convenient set of assertion matchers in tests
* [Gomicro Cucumber Docker Image](https://github.com/gomicro/docker-cucumber) - used to execute BDD style tests written in cucumber
* [Gomicro Service](https://github.com/gomicro/service) - used to bootstrap service (dockerfile, docker-compose, minimal running go service)
* [SQLMock](https://github.com/DATA-DOG/go-sqlmock) - used for mocking the database and enable unit testing of database calls
* [UUID](https://github.com/google/uuid) - used to facilitate uuid parsing

## Notes

### Decisions

The geoIP lookup provides a radius as part of its results. Since a person could have performed a login event from anywhere within that radius around the provided location, the radius was used to reduce the minimum possible distance a person could have to travel between two login events. This was done in order to favor reducing false positives and only result in flagging suspicious activity that has a significantly higher chance of being truly suspicious.

### Deployment

Minor code changes would be needed to switch the service from using SQLite to a more production ready database such as Postgres. Mostly this would consist of the configuration needing to be updated, and the bootstrapping of the database could be moved out of the service and into a migration tool or other form of management independent of the service.

Deployment would involve pushing the resulting docker image from `make build` to a docker image repository and then deploying that image as a service into a container orchestrator such as ECS (Elastic Container Service) or Kubernetes.

### Testing

While a good set of checks can be provided by unit tests, not everything is best covered by them. It is possible to setup endpoint handlers in Go to be covered by unit tests, but the effort of setting up the mocks of a test server, database, and more start to outweigh the value of black box testing. The white box type of tests as just mentioned majorly only provide value in testing states that are more rare occurrences and with idiomatic Go error handling, are at least handled. As such, I have chosen to address coverage of the endpoints with BDD style tests which do a better job of articulating the desired feature in user centric terms rather than highly technical descriptions. This drives better empathy for the end user and increases communication with more universal terms that non developers can provide input against.


### SQLite Requirement

Normally I would not use SQLite even for local development, as a significantly closer to deployed code setup can be accomplished with a Postgres container.

Specific Items of Note:

* SQLite requires `CGO` to be enabled and requires either installing build tools or using the non `alpine` Golang image
* The build on a fresh build instance takes longer due to the size difference of the `golang-alpine` vs `golang` docker images (370MB vs 809MB)
* The resulting docker image is also larger, because we cannot run a scratch container (25MB vs 99MB)
* Bootstrapping of the database gets built into the binary rather than being able to build it into the database image or use some migration tool that would apply to prod
