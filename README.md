## Combine Golang (echo) and ReactJS into single docker container
This is project will show you how to combine ReactJS and Golang into single container.

## Usage
In the project I set default port is `8080`. You can change it by set `$PORT` in your environment.
### Prerequisite
* **With docker:** You need installed `docker` *(if you are run it with docker)* and `docker-compose` *(if you run it with docker-compose)*. I placed links tutorial in here ([docker](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04) and [docker-compose](https://docs.docker.com/compose/install/)) 
* **Without docker:** You need install `go (>=1.14)`, `nodejs (>= 12)` and `yarn`.  
### With `docker`
```shell
docker build -t <you_choose_image_name>:<tag> .
docker run -d -p 8080:8080 <you_choose_image_name>:<tag> 
```

### With `docker-compose`
```shell
docker-compose up --build
```

### Without `docker`
```shell
# Build need to build `frontend` first
cd ./frontend
yarn build
# After this phase you will see a build folder inside current folder
cd ..
go mod download
go build -o app . && ./app
```

