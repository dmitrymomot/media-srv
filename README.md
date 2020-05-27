# media-srv
Media Service API for image processing.

## Prerequisites
- You need installed docker and docker-compose machine, where the app will be runned
- You need one of following cloud storage:
  - AWS S3
  - Digitalocean Staces

## Run the app

### Clone the repo
```
git clone github.com/dmitrymomot/media-srv && cd media-srv
```

### Set up your cloud storage provider config
```
cp .env.local .env
```

### Run the app
```
make build docker up
```
The API will be served on port `:8080`, you can change it in `.env` file

### Stop and remove all containers
```
make downn
```

> Run `make` or `make help` to see all available commands


## Documentation
It doesnn't exist, but you can use [Insomnia app]([http](https://insomnia.rest)) config file (grab it from the root of this repo) to see examples of all requests and responses.