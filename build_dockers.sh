#!/bin/sh

cd button_clicker/
docker build -t button-clicker .
docker run -d -p 8080:8080 --name button_container button-clicker
cd ../

cd hidden_flag/
docker build -t hidden-flag .
docker run -d -p 8020:8020 --name hidden_flag_container hidden-flag
cd ../

cd posts/
docker build -t posts .
docker run -d -p 8010:8010 --name posts_container posts
cd ../
