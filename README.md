# Persons Service

## Overview

This is a demo implementing sqlite and unit tests.

## Local setup

### Docker

The easiest way to run this application is to build a docker image from it and run that image as a container. This would
handle everything from compiling the executable to generating the swagger documentation without much need to know how
any of that works.

To do that run

```
docker build -t persons-service:latest .
```

and then

```
docker run --name persons-service -p 8080:8080 persons-service-service:latest
```

in the project's root directory.

The downside to using this method of compiling and executing the app is that it makes debugging a little more
complicated. If you're using VSCode for development, you can find information on how to get that right over here

- https://code.visualstudio.com/docs/containers/debug-common.

### Non-Docker setup

The non-docker method of compiling and running this project involves having to run a few commands to test/compile the
code as well as generate swagger documentation. Usually you would have to run these commands sequentially in your
terminal, but we've built in a shortcut with the help of a tool called `modd`.

Not only does modd save us from having to constantly enter a handful of commands every time we'd like to re-compile and
run our project, it also serves as a hot reloader. So (in most cases) it will re-compile and serve your code changes on
the fly, without you having to stop and start things up again.

If you feel the project compilation and startup could use some tweaking, you would need to make those changes
in `${projectRoot}/modd.conf`. Just make sure that, if ever you have to do this, you __make those same changes to the
Dockerfile__ also, or what you have in developement won't match the other dockerised environments (test/staging/prod).

#### Setup steps:

1. Install modd: https://github.com/cortesi/modd.git
2. Install go-swagger: https://github.com/go-swagger/go-swagger.git
3. In `${projectRoot}`, create a `settings.yml`, containing the following:

    ```yaml
    service:
        name: "persons-service"
        path_prefix: ""
        version: "1"
        debug: false
        port: 8080
    database:
        engine: "sqlite"
        host: ""
        name: "mytest"
        port: 3306
        user: ""
        password: ""
    ```

4. Run the app by simply running the command:

    ```
    modd
    ```

   To run the app without modd pre-compiling and hot reloading, run:

    ```
    go run main.go
    ```

   To generate swagger docs without modd or docker:

    ```
    swagger generate spec -o internal/api/v1/swagger/swagger.yml
    ```

5. Try out some of the test endpoints

   Download an API client like Postman (https://www.postman.com/downloads/) to be able to test the project API.

   Try the health check endpoint (GET localhost:8080/healthcheck).

## Project architecture

`/settings`

*️ This is where everything related to project configuration and settings is kept. Settings can be read from yaml files
and can be stored in the same package.

`/files`

*️ Files holds all static files to be served. This includes the generated swagger documentation.

`/internal`

*️ This is the main package of our project where everything specifically related to its domain and functionality lives.

`/internal/api`

*️ This package contains everything related to the service api - router, routes, handlers, middleware(filters) etc.

`/internal/api/${versionNumber}/swagger`

*️ Holds generated swagger files

`/internal/persons/service`

*️ This package contains all business logic and data repository interaction for persons-service domain.

`/internal/persons/entity`

*️ This package contains structs matching the raw data fetched from each repository. This data will then might be
transformed and enriched in the api layer before it is emitted.

`/internal/persons/persons.go`

*️ This file serves as the contract (interface) for functionality in the packages within, namely repository and service.

`/internal/persons/repository`

*️ This package contains all implementations of the `Repository` interface declared
in `/internal/persons-service/persons/persons.go`.
