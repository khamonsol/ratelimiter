#!/bin/zsh
docker build -t ratelimiter:latest .

docker stop ratelimiter-nginx
docker rm ratelimiter-nginx
docker run --name ratelimiter-nginx -d -p 1337:80 ratelimiter:latest
docker logs ratelimiter-nginx
