# Stability & Reliability

Sample app taken from [demo-go-postgres repo](https://github.com/rsmcode/demo-go-postgres).

## Application Containerization

### Docker

Builds Docker images from a Dockerfile and a “context”. A build’s context is the set of files located in the specified PATH or URL. The build process can refer to any of the files in the context. For example, your build can use a COPY instruction to reference a file in the context.

### Kaniko

kaniko is a tool to build container images from a Dockerfile, inside a container or Kubernetes cluster.

kaniko doesn't depend on a Docker daemon and executes each command within a Dockerfile completely in userspace. This enables building container images in environments that can't easily or securely run a Docker daemon, such as a standard Kubernetes cluster.

### Buildpack

A buildpack is something you’ve probably leveraged without knowing, as they’re currently being used in many cloud platforms. A buildpack’s job is to gather everything your app needs to build and run, and it often does this job quickly and quietly.

That said, while buildpacks are often a behind-the-scenes detail, they are at the heart of transforming your source code into a runnable app image.

Auto-detection 
What enables buildpacks to go unnoticed is auto-detection. This happens when a platform sequentially tests groups of buildpacks against your app’s source code. The first group that deems itself fit for your source code will become the selected set of buildpacks for your app. Detection criteria is specific to each buildpack – for instance, an NPM buildpack might look for a package.json, and a Go buildpack might look for Go source files.

### KO

ko is a simple, fast container image builder for Go applications.

It's ideal for use cases where your image contains a single Go application without any/many dependencies on the OS base image (e.g., no cgo, no OS package dependencies).

ko builds images by effectively executing go build on your local machine, and as such doesn't require docker to be installed. This can make it a good fit for lightweight CI/CD use cases.

ko is an open-source tool developed at Google that helps you build container images from Go programs and push them to container registries (including Container Registry and Artifact Registry). ko does its job without requiring you to write a Dockerfile or even install Docker itself on your machine.

ko is spun off of the go-containerregistry library, which helps you interact with container registries and images. This is for a good reason: The majority of ko’s functionality is implemented using this Go module. Most notably this is what ko does:

Download a base image from a container registry
Statically compile your Go binary
Create a new container image layer with the Go binary
Append that layer to the base image to create a new image
Push the new image to the remote container registry

For a deeper comparision of KO and Buildpack, checkout the [Google Cloud Blog](https://cloud.google.com/blog/topics/developers-practitioners/ship-your-go-applications-faster-cloud-run-ko)



## Local Testing

### KIND (Kubernetes in Docker)

kind is a tool for running local Kubernetes clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.

## Manifest Packaging

### Helm

Helm uses a packaging format called charts. A chart is a collection of files that describe a related set of Kubernetes resources. A single chart might be used to deploy something simple, like a memcached pod, or something complex, like a full web app stack with HTTP servers, databases, caches, and so on.

Charts are created as files laid out in a particular directory tree. They can be packaged into versioned archives to be deployed.


### Tanka

Grafana Tanka is the robust configuration utility for your Kubernetes cluster, powered by the unique Jsonnet language

### KO

"ko resolve" and "ko apply" commands you can hydrate your YAML manifests as ko replaces your "image:" references in YAML automatically with the image it builds, so you can deploy the resulting YAML to the Kubernetes cluster with kubectl:

## Deploy

### ArgoCD

Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes. Application definitions, configurations, and environments should be declarative and version controlled. Application deployment and lifecycle management should be automated, auditable, and easy to understand. Argo CD follows the GitOps pattern of using Git repositories as the source of truth for defining the desired application state. Kubernetes manifests can be specified in several ways:

## Devex

### Octant

Octant is a tool for developers to understand how applications run on a Kubernetes cluster. It aims to be part of the developer's toolkit for gaining insight and approaching complexity found in Kubernetes. Octant offers a combination of introspective tooling, cluster navigation, and object management along with a plugin system to further extend its capabilities.