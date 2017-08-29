[![](https://images.microbadger.com/badges/image/infinityworks/moby-container-stats.svg)](http://microbadger.com/images/infinityworks/moby-container-stats "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/infinityworks/moby-container-stats.svg)](http://microbadger.com/images/infinityworks/moby-container-stats "Get your own version badge on microbadger.com")

# Moby Container Stats

Exposes basic metrics for your containers directly from the moby stats API to a prometheus compatible endpoint. Doing a similar job to the popular cAdvisor application, this is a much more lightweight and simplified design for just the key metrics that people require. It requires less resources to run as an exporter on your hosts, it's also less expensive to store in prometheus.

## Configuration

This exporter is setup to take input from environment variables:

### Required
None..

### Optional
* `LISTEN_PORT` The port you wish to run the container on, the Dockerfile defaults this to `9244`
* `METRICS_PATH` the metrics URL path you wish to use, defaults to `/metrics`
* `LOG_LEVEL` The level of logging the exporter will run with, defaults to `debug`


## Install and deploy

Run manually from Docker Hub:
```
docker run -d --restart=always -p 9244:9244 infinityworks/moby-container-stats
```

Build a docker image:
```
docker build -t <image-name> .
docker run -d --restart=always -p 9244:9244 <image-name>
```

## Docker compose

```
github-exporter:
    tty: true
    stdin_open: true
    ports:
      - 9244:9244
    image: infinityworks/moby-container-stats:latest
```

## Metrics

Metrics will be made available on port 9244 by default
An example of these metrics can be found in the `METRICS.md` markdown file in the root of this repository