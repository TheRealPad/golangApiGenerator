![Unit test](https://github.com/TheRealPad/golangApiTemplate/actions/workflows/unit-test.yml/badge.svg)
![Integration test](https://github.com/TheRealPad/golangApiTemplate/actions/workflows/integration-test.yml/badge.svg)
![Deploy](https://github.com/TheRealPad/golangApiTemplate/actions/workflows/deploy.yml/badge.svg)

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
- /health               -> GET
- /health/html          -> GET
- /health/traffic       -> GET
- /health/traffic/html  -> GET
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