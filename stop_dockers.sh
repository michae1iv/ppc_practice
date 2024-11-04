#!/bin/sh

docker stop button_container
docker rm button_container

docker stop hidden_flag_container
docker rm hidden_flag_container

docker stop posts_container
docker rm posts_container
