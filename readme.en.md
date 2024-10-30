# Alioth App Template

![Go Version](https://img.shields.io/github/go-mod/go-version/alioth-center/alioth-app-template)
![Go Report Card](https://goreportcard.com/badge/github.com/alioth-center/alioth-app-template)
![GitHub Actions](https://img.shields.io/github/actions/workflow/status/alioth-center/alioth-app-template/build-docker.yml?branch=main)
![License](https://img.shields.io/github/license/alioth-center/alioth-app-template)

[中文](./readme.md) | **English**

----

## Features

1. A web service template based on the Gin framework
2. Pre-configured project architecture, including configuration files, logging, databases, etc.
3. Includes an example endpoint and implementation
4. Configured Dockerfile and Github Actions for automatic builds

## Usage

1. Fork this project
2. Change the project name and module path, edit the configuration file
3. Start developing your business logic
4. Run the project
5. Debug and test your endpoints
6. Push your code to **your** repository
7. Enable Github Actions for automated builds
8. Wait for the build to complete and check the build results
9. Use your Docker image and deploy your service

## Directory Structure

```text
.
├── app
│   ├── api
│   ├── entity
│   ├── global
│   ├── router
│   └── service
├── config
├── go.mod
├── go.sum
├── main.go
└── readme.md
```

- `app` Business logic code
    - `api` API definitions
    - `entity` Operational entities, such as request and response structures
    - `model` Data model (optional)
    - `dao` Data access layer (optional)
    - `global` Global variables
    - `router` Route configurations
    - `service` Business logic
- `config` Configuration files
- `main.go` Project entry point

## Default Endpoint

The template includes an `example/ping` endpoint to test if the service is running correctly.

Here’s how to send a request to this endpoint. It returns the content of your request as-is.

```shell
curl -X POST http://localhost:10000/example/ping \
  --header "Content-Type: application/json" \
  --data-raw '{"hello": "world"}'
```

## Contributors

![Contributors](https://contrib.rocks/image?repo=alioth-center/alioth-app-template&max=1000)
