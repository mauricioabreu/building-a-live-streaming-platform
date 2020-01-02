# Building a live streaming platform

Code for the [Building a live streaming platform](https://www.maugzoide.com/posts/building-a-live-streaming-platform-part-i/) post

# Requisites

**Before running these commands you need to create a Docker network to link containers**
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