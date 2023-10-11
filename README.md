Example project with Go and Benthos Prometheus metrics producers, to test how Prometheus reads metrics.

## Run Go client
```
go run main.go
```

View metrics at http://localhost:2112/metrics


## Run Benthos pipeline
```
benthos -c config.yaml
```

## Run Prometheus server
```
docker run \
    -p 2113:9090 \
    -v $PWD/prometheus.yaml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```

View Prometheus at http://localhost:2113