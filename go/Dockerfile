FROM golang:latest as dev
RUN apt-get update \
 && apt-get install -y vim \
 && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# RUN go run main.go