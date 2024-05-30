![Unit test](https://github.com/TheRealPad/golangApiTemplate/actions/workflows/unit-test.yml/badge.svg)
![Integration test](https://github.com/TheRealPad/golangApiTemplate/actions/workflows/integration-test.yml/badge.svg)
![Deploy](https://github.com/TheRealPad/golangApiTemplate/actions/workflows/deploy.yml/badge.svg)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

# Go API from config file

Create a server where you just need a config file, and the program do the rest

## Config
- server name
- api port
- routes (name and HTTP method + describe which part of CRUD it is)
- data models (name, fields and fields type)

## How to run

By default, the API is just a basic API with a logger to get the activity on the software

Default available routes:
```txt
- /health               -> GET # display information about start time of the api, number of api calls and data models
- /health/html          -> GET # display above information in html syntax
- /health/traffic       -> GET # show all api calls
- /health/traffic/html  -> GET # display above information in html syntax
```

If you're on Unix-like system, you can use the [Makefile](Makefile) to run the Docker in background

If you can't use the Makefile, you can run:
```bash
docker-compose -f ./config/docker/docker-compose.yml up --build -d
```

If you can't use Docker, you can run:
```bash
go run .
```

## Database

For now the API only handle MongoDB, you can create a cluster and pass an url looking like that:
```
mongodb+srv://username:password@cluster0.k1vyunp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
```

This is the tutorial I followed to get the url https://www.mongodb.com/docs/drivers/go/current/quick-start/