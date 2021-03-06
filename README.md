# Building a live streaming platform

## Part I

Code for the [Building a live streaming platform](https://www.maugzoide.com/posts/building-a-live-streaming-platform-part-i/) post

### Requisites

**Before running "Usage" commands you need to create a Docker network to link containers**
```
docker network create live-stream
```

You may need *sudo* to run it as *root*.

### Usage

To run any command listed here, make sure you have **make** and **Docker** installed.

Run a container with NGINX-RTMP
```
make runserver
```

Ingest a live video using ffmpeg
```
make ingest
```

## Part II

### Web Application

The web application is used to authorize the incoming streams. Is there a streamer you like? Some famous person you like the videos? I watch some game streamers. How does a video platform know that your preffered Youtuber is himself/herself?

Before starting our web application, we need to create the database to hold our *publisher credentials*:

```
make create-db
```

And then we can use [docker-compose](https://docs.docker.com/compose/) to start our live streaming service:

```
make run-live
```

Now we have a server running. We can send HTTP requests to the server to experiment it.

### API

We can test our API using *cURL*, making some requests and reading the response body along with our application logs.

#### `/auth` (simple - with name and key)

```shell
curl -XPOST -H "Accept: application/json" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:9090/auth -d "name=foos&psk=bar"
```

#### `/auth` (improved - with key only)

```shell
curl -XPOST -H "Accept: application/json" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:9090/auth -d "name=bar"
```
