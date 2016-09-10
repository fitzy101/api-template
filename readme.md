## Running API with docker.
You can run the API using a docker container. You will need to pass a number of environmental variables to the container for it to run.

```
#!shell
cd $GOPATH/github.com/fitzy101/api-template

docker build -t api-template .

docker run -p 8090:3000 -d \
--env DB_USERNAME=root \
--env DB_PASSWORD=password \
--env DB_HOST=localhost \
--env DB_PORT=3306 \
--env DB_SCHEMA=default \
api-template

curl localhost:8090/status
$ 
```

