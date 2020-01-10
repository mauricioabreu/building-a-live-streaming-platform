# Building a live streaming platform

Code for the [Building a live streaming platform](https://www.maugzoide.com/posts/building-a-live-streaming-platform-part-i/) post

## Requisites

**Before running "Usage" commands you need to create a Docker network to link containers**
```
docker network create live-stream
```

You may need *sudo* to run it as *root*.

## Usage

To run any command listed here, make sure you have **make** and **Docker** installed.

Run a container with NGINX-RTMP
```
make runserver
```

Ingest a live video using ffmpeg
```
make ingest
```

## API

We can test our API using *cURL*, making some requests and reading the response body along with out application logs.

### `/auth`

```shell
curl -XPOST -H "Accept: application/json" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8080/auth -d "name=foos&psk=bar"
```