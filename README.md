# Superman

## Building
All steps below are run the same way they are in the build. The build may chain the steps slightly differently for the benefit of reporting, but otherwise does the same thing.

### Running

The service will be stood up in a local docker environment and bound to `localhost:4567` for allowing any manual poking and prodding of the service.

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

## Notes

### SQLite Requirement

Normally I would not use SQLite even for local development, as a significantly closer to deployed code setup can be accomplished with a postgres container.

Specific Items of Note:

* SQLite requires `CGO` to be enabled and requires either installing build tools or using the non `alpine` Golang image
* The build on a fresh build instance takes longer due to the size difference of the `golang-alpine` vs `golang` docker images (370MB vs 809MB)
* The resulting docker image is also larger, because we cannot run a scratch container (25MB vs 99MB)
* Bootstrapping of the database gets built into the binary rather than being able to build it into the database image or use some migration tool that would apply to prod
