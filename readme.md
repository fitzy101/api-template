# api-template
A Golang API template that can be used for anything.

To build the docker container, run:

```sh
cd $GOPATH/src/github.com/fitzy101/api-template
docker build -t api-template .
```

You can then run the API using the command (update the database creds/info 
accordingly):
```sh
docker run -p 8090:3000 -d \
--env DB_USERNAME=root \
--env DB_PASSWORD=password \
--env DB_HOST=127.0.0.1 \
--env DB_PORT=3306 \
--env DB_SCHEMA=default \
api-template
```

Confirm the API is running with:
```sh
curl localhost:8090/status
```
