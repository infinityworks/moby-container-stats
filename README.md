[![Go Report Card](https://goreportcard.com/badge/github.com/infinityworks/moby-container-stats)](https://goreportcard.com/report/github.com/infinityworks/moby-container-stats)[![](https://images.microbadger.com/badges/image/infinityworks/moby-container-stats.svg)](http://microbadger.com/images/infinityworks/moby-container-stats "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/infinityworks/moby-container-stats.svg)](http://microbadger.com/images/infinityworks/moby-container-stats "Get your own version badge on microbadger.com")

# Moby Container Stats

Exposes basic metrics for your containers directly from the moby stats API to a prometheus compatible endpoint. Doing a similar job to the popular cAdvisor application, this is a much more lightweight and simplified design for just the key metrics that people require. It requires less resources to run as an exporter on your hosts, it's also less expensive to store in prometheus. Metrics are obtained through async calls to the moby stats API over the Docker socket. Scrapes can sometimes take up to 2 seconds, this is due to the inherent way moby calculates it's CPU usage at present. Ideally, the work Docker are doing towards native prometheus metrics for containers will render this exporter obsolete.

## Configuration

This exporter is setup to take input from environment variables, please make sure you also map in the docker socket to the container, details below.

### Required
None..

### Optional
* `LISTEN_PORT` The port you wish to run the container on, the Dockerfile defaults this to `9244`
* `METRICS_PATH` the metrics URL path you wish to use, defaults to `/metrics`
* `LOG_LEVEL` The level of logging the exporter will run with, defaults to `debug`


## Install and deploy

Run manually from Docker Hub:
```
docker run -d --restart=unless-stopped -v /var/run/docker.sock:/var/run/docker.sock -p 9244:9244 infinityworks/moby-container-stats
```

Build a docker image:
```
docker build -t <image-name> .
docker run -d --restart=unless-stopped -v /var/run/docker.sock:/var/run/docker.sock -p 9244:9244 <image-name>
```

## Docker compose

```
moby-container-exporter:
    tty: true
    stdin_open: true
    ports:
      - 9244:9244
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    image: infinityworks/moby-container-stats:latest
```

## Metrics

Metrics will be made available on port 9244 by default
An example of these metrics can be found in the `METRICS.md` markdown file in the root of this repository
