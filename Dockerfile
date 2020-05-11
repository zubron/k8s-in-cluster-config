FROM golang:1.14-buster as build

COPY . /k8s-in-cluster-config
WORKDIR /k8s-in-cluster-config

RUN go build

FROM debian:buster-slim

COPY --from=build /k8s-in-cluster-config/k8s-in-cluster-config /k8s-in-cluster-config
CMD ["bash", "-c", "/k8s-in-cluster-config"]
