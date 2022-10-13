# HA Challenge

## Project
 
|2218. Space exploration is underway and mostly done by private companies. You joined
Engineering department of one of the private government contractors, Atlas Corporation.
In that year and age, everything is automated, so survey and data gathering is done by drones.
Unfortunately, although drones are perfectly capable of gathering the data, they have issues
with locating databank to upload gathered data. You, as the most promising recruit of Atlas
Corp, were tasked to develop a drone navigation service (DNS):
- each observed sector of the galaxy has unique numeric SectorID assigned to it
- each sector will have at least one DNS deployed
- each sector has different number of drones deployed at any given moment
- it’s future, but not that far, so drones will still use JSON REST API

## Docker Environment

### Dockerfile

The Dockerfile will build the project and the image carry the service itself

### Docker-compose

There's a docker-compose.yml file included you can use as a starting point.
You can configure the service with environment variables.

The project should work out of the box using
```
docker-compose up --build
```

## Configuration

The service is configurable with environment variables.

- HA_ADDRESS: Listen address. Default is :8080
- HA_SECTOR_ID: The specified sector id. It is mandatory parameter. It must be integer.

### Example
```
HA_ADDRESS=:8080
HA_SECTOR_ID=1
```

## Code structure of the service

This service has 3 main component.
- cmd: this is the runnable package
- service: these calculate the locations
- web: this build up the webserver and the controllers

## API doc

The API has been documented in Apiary format. Please check the *apidoc.apib* file for more details