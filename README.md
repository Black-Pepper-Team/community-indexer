# community-indexer

NOTE: This service executes part of the logic of the mobile application such as ZK proofs generation and smart contracts 
calls to decrease the mobile developer's load and deliver the working Proof Of Concept in 2 days. So don't be severe
about code quality and secret key handling :D.

## About

Here you can find the logic of the ZKPs generation to register in the AuthStorage contract and then authenticate on the
Chat contract to leave an anonymous message.

## Install

  ```
  git clone github.com/black-pepper-team/community-indexer
  cd community-indexer
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main migrate up
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.  

## Running from docker 
  
Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

  ```
  docker build -t github.com/black-pepper-team/community-indexer .
  docker run -e KV_VIPER_FILE=/config.yaml github.com/black-pepper-team/community-indexer
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `migrate up` command to create database schema
* Launch the service with `run service` command


### Database
For services, we do use ***PostgresSQL*** database. 
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).
