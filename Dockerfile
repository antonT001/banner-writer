FROM golang:latest
RUN apt-get update && apt-get upgrade -y
RUN apt-get install git faketime -y