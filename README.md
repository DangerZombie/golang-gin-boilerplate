# Golang Gin Boilerplate
This is the boilerplate of Golang using Gin Framework

## 1. Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### 1.1. Prerequisites

These are the prerequisite library to run the project:

- [Go](https://golang.org/doc/install)
- [mockgen](https://github.com/uber-go/mock)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [GNU Make](https://www.gnu.org/software/make/)
<!-- - [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [wire](https://github.com/google/wire)
- [migrate](https://github.com/golang-migrate/migrate)
- [gcloud CLI](https://cloud.google.com/sdk/docs/install)
- [gentool](https://gorm.io/gen/gen_tool.html) -->

### 1.2. Installing

A step-by-step series of examples that tell you how to get a development env running

1. Clone the repository

```bash
git clone
```

2. Intialize the project

```bash
make init
```

3. Open the cloned repository. We recommend you to use Visual Studio Code because it's free and light but powerfull. [Visual Studio Code](https://code.visualstudio.com/download)

4. Create lauch.json, to make you eaiser to run and debug the program. This step if you use Visual Studio Code.

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch file",
            "type": "go",
            "request": "launch",
            "env": {
                "localhost": "8080" // this is the default port, you can use your preferrable port
            },
            "mode": "debug",
            "program": "main.go"
        }
    ]
}
```

5. Ask the Backend team on the latest `config-*.yaml` for running the project. Please note this config will not be uploaded to the repository.

## 2. Development

### 2.1. Working With Tests

This section describes how to work with tests in this repository.

### 2.1.1. Writing the tests

#### 2.1.1.1. Writing the unit tests

For every `.go` source code, write a corresponding `_test.go` file in the same directory. For example, if you have a `foo.go` file in `service/foo`, then you should have a `foo_test.go` file in the same directory. Just a quick reminder that unit tests do not have any dependencies to external services such as database, cache, etc. We use `gomock` to mock the dependencies.

#### 2.1.1.2. Writing the integration tests

This is similar with the unit tests. For every `.go` source code, write a corresponding `_integration_test.go` file in the same directory. For example, if you have a `foo.go` file in `repository/foo`, then you should have a `foo_integration_test.go` file in the same directory.

The test function has to use prefix `TestIntegration_` so that we can run the integration tests separately from the unit tests. After the prefix, you can name the test function according to the test case or function that you want to test. Right below the function name, you have to put the following line:

```go
if testing.Short() {
    t.Skip("skipping integration test")
}
```

This is to skip the integration test when we run the unit tests.

### 2.1.2. Running the tests

Run the following to run the tests

```bash
make generate_mocks
```

Excute this to run the unit test

```bash
make test_unit
```

Execute this to run the integration test
```bash
make test_integration
```

### 2.2. Working with Docker

This project is using Docker Container for easier development process. This container could replicate the server environment, while you use build and run the project correctly on Docker, it will be also run correctly in your server.

We suggest that you use the Docker CLI, during this process. After installing the `Docker` and `Docker Compose` you can start the virtualization. This step will produce in Linux/WSL.

### 2.2.1 Installing the Docker

1. Update package and install dependency

```bash
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
```

2. Add Docker GPG Key and Repository

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

3. Update package and install Docker

```bash
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io
```

4. Run Docker without sudo ($USER could change with you user on your device)

```bash
sudo usermod -aG docker $USER
```

5. Restart your terminal, then check the Docker version

```bash
docker --version
```

### 2.2.2 Installing the Docker Compose

You also can use `Docker Compose` if you want to simplify the environment setup process.

1. Download Docker Compose

```bash
sudo apt update
sudo apt install docker-compose
```

2. Check the Docker Compose version

```bash
docker-compose --version
```

### 2.2.3 Running with Docker

This section will run the application on Docker container and you will need to provide the database and other application.

1. Prepare the Dockerfile and build the Docker image

```bash
docker build -t golang-gin-boilerplate:1.0 .
```

2. Run the Docker image

```bash
docker run -d -p 8080:8080 golang-gin-boilerplate:1.0
```

3. Check the running container

```bash
docker ps -a
```

### 2.2.4 Running with Docker Compose

If you use the Docker Compose you also can dockerize the database or the other application, Remember to make this service running in the same network [link](https://docs.docker.com/network/).

1. Run the `docker-compose.yaml`, make sure your terminal path in the root path of the `docker-compose.yaml`

```bash
docker-compose up
```

2. Check the running image

```bash
docker ps -a
```

### 2.2.5 Other Docker command

Terminate the running container

```bash
docker stop $CONTAINER_ID
```

Remove Docker Container

```bash
dokcer rm $CONTAINER_ID
```

Restart Docker

```bash
sudo service docker restart
```

Check Docker Status
```bash
sudo service docker status
```

### 2.3 Swagger Documentation

For easier test the API create the swagger command in your code. This comment is representing the OpenAPI structure.

Download the swagger file first on [go-swagger](https://github.com/go-swagger/go-swagger/releases) repository. Then copy to your `usr/local/bin`.

Generate Swagger documentation

```bash
make generate_swagger
```

Read-only documentation page `http://localhost:8080/doc`
Attractive documentation page `http://localhost:8080/docs`

### 2.4. Working Directory

This project also have many folder inside, then to minimize the complexity the directory will as follow:

```
project/
├── .github/
│   ├── workflow/
│   └── ...
├── endpoint/
│   └── ...
├── helper/
│   ├── auth/
│   ├── database/
│   ├── http_helper/
│   ├── message/
│   ├── static/
│   ├── util/
│   └── ...
├── http/
│   └── ...
├── initialization/
│   └── ...
├── model/
│   ├── base/
│   ├── entity/
│   ├── request/
│   ├── response/
│   └── ...
├── repository/
│   ├── repository_user/
│   └── ...
├── service/
│   ├── service_user/
│   └── ...
```

Explanations:

1. `.github` folder contains whole files related github.
2. `endpoint` folder contains whole files for related the endpoint of the API.
3. `helper` folder contains whole files for help code easier such as auth logic, conversion logic, etc.
4. `http` folder contains files conversion result after processing and convert according with the response.
5. `model` folder contains model entity represent of the database column.
6. `repository` folder contains files to querying the data to database.
7. `service` folder contains files with business logic or other logic.