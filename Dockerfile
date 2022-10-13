FROM golang:1.18
WORKDIR /mnt/work
COPY . .
RUN go build -o ha-dns ./cmd

# Docker is used as a base image so you can easily start playing around in the container using the Docker command line client.
FROM docker
COPY --from=0 /mnt/work/ha-dns /usr/local/bin/ha-dns