# Implementing Microservices & Kubernetes in Go

This repository contains some quick work that I did during the beginning of my
internship at [Ori](https://ori.co) during secondary school. These microservice
projects merely served as a way to get up to speed in Golang, so that I could
engage further with their codebase. The areas covered in these repos are:

* gRPC

* Protocol Buffers

* Minikube & Kubernetes

The contained directories are as follows:

* [godomicroservice](godomicroservice): As the read-me in the subdirectory
  explains, this is the runner of a simple cloud native microservice written in
  Go, using gRPC and Protocol Buffers.

* [gotrymicroservice](gotrymicroservice): Wrapper around [godomicroservice](godomicroservice)
  to run it in a mini Kubernetes cluster (using local [Minikube](https://minikube.sigs.k8s.io/docs/)).

* [godotest](godotest): Testing the above microservices using native Go tests
  and the [Cirrus](https://cirrus-ci.org/) continuous integration (CI) tool
  suite.
