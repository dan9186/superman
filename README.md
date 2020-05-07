# Superman

## Building
All steps below are run the same way they are in the build. The build may chain the steps slightly differently for the benefit of reporting, but otherwise does the same thing.

### Running

The service will be stood up in a local docker environment and bound to `localhost:4567` for allowing any manual poking and proding of the service.

```
make run
```

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
