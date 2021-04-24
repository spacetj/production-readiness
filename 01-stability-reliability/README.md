# Stability & Reliability

- Blog: [CloudNative Production Readiness Part 1](https://tejasc.com/cnpr-part1/)
- Sample App: [demo-go-postgres repo](https://github.com/rsmcode/demo-go-postgres)

The following cloud native tools helps:

- build and package apps in a standardized and automated fashion
- enable standardize development, testing and debugging cycles

## Prerequisite

```bash
make deps
```

> Only execute this command if you're on a mac and you have brew installed

The command will install the following tools:

- GolangCI-Lint
- Kind
- Skaffold
- KO
- Tanka
- Helm
- Pack
- ArgoCD
- Octant

## Application Containerization

### Docker

<img src="./assets/docker.png" width="100px" height="80px" align="left" style="padding-right:15px"/>

Docker is a tool for building applications based on containers — lightweight execution environments that shares the operating system kernel but otherwise runs in isolation from one another. Utilises Dockerfile, context and Docker Daemon to build the image.

| Open Source                         | Maintained By | First Release |
| ----------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/docker) ✅ | Docker        | 20 March 2013 |

```bash
make build.docker
```

The command builds a docker image using the instructions defined in the [Dockerfile](./Dockerfile). It creates a local image called `todo:docker`.

---

### Kaniko

<img src="./assets/kaniko.png" width="100px" height="100px" align="left" style="padding-right:15px"/>

Kaniko is a tool for building container images from a Dockerfile without requiring a Docker daemon. Kaniko executes each command in userspace, enabling building container images in environments that can't easily or securely run a Docker daemon, such as a Kubernetes cluster. Kaniko can also be used to speed up builds by caching image layers.

| Open Source                                              | Maintained By | First Release |
| -------------------------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/GoogleContainerTools/kaniko) ✅ | Google        | 18 May 2018   |

```bash
make build.kaniko
```

Builds a docker image using kaniko and the instructions defined in the [Dockerfile](./Dockerfile). Kaniko requires the executor container (gcr.io/kaniko-project/executor) to build the image. Images are built and pushed directly to gcr. Layers are cached by executing kaniko with the following flags: `--cache=true --cache-repo "${GCR_IMAGE}/cache"`.

---

### Buildpack

<img src="./assets/buildpack.png" width="100px" height="100px" align="left" style="padding-right:15px"/>

Buildpack can build images without a Dockerfile. It uses auto-detection instead. It tests groups of buildpacks against your source code, and the first group that fits your source code will become the selected set of buildpacks. Criteria are specific to each buildpack – for instance, an NPM buildpack looks for a package.json, and a Go buildpack looks for Go source files, etc.

| Open Source                                  | Maintained By | First Release |
| -------------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/buildpacks/pack) ✅ | CNCF          | 21 Aug 2018   |

```bash
make build.pack
```

The command uses Google Buildpack (gcr.io/buildpacks/builder:v1) to create a local docker image called `todo:pack`.

---

### ko

<img src="./assets/ko.png" width="160px" height="80px" align="left" style="padding-right:15px"/>

ko builds images by effectively executing go build on your local machine, and as such doesn't require a Dockerfile or docker to be installed. Ideal for use cases where your image contains a single Go application without any dependencies on the OS base image (e.g., no cgo, no OS package dependencies).

ko does the following:

- Download a base image from a container registry
- Statically compile your Go binary
- Create a new container image layer with the Go binary
- Append that layer to the base image to create a new image
- Push the new image to the remote container registry

| Open Source                            | Maintained By | First Release |
| -------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/google/ko) ✅ | Google        | 22 Mar 2019   |

```bash
make build.ko
```

The command uses `ko` to create and push docker image to gcr registry.

---

## Local Testing

### KIND (Kubernetes in Docker)

<img src="./assets/kind.png" width="100px" height="65px" align="left" style="padding-right:15px"/>

kind is a tool for running local Kubernetes clusters using Docker container nodes. kind was primarily designed for testing Kubernetes itself, but can be used for local development or CI.

| Open Source                                       | Maintained By | First Release |
| ------------------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/kubernetes-sigs/kind) ✅ | CNCF          | 29 Nov 2018   |

```bash
make cluster
```

The command uses that to create a kubernetes cluster. Any images you want to deploy to kind needs to be manually loaded into kind using the following command: `kind load docker-image my-custom-image-0 my-custom-image-1`

---

## Manifest Packaging

### Helm

<img src="./assets/helm.svg" width="100px" height="100px" align="left" style="padding-right:15px"/>

Helm uses charts, a collection of files that describe a related set of Kubernetes resources. Charts are created as files laid out in a particular directory tree. They can be packaged into versioned archives to be deployed. Helm uses Sprig template library to help template the yaml files and values can be defined in single or multiple value files.

| Open Source                            | Maintained By | First Release |
| -------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/helm/helm) ✅ | CNCF          | November 2015 |

```bash
make manifests.helm
```

Uses the [helm chart](./deploy/helm) to template files and outputs them [here](./deploy/output/helm.yaml). You can define default values in the [chart values](./deploy/helm/values.yaml) file and override values on per env basis which is defined in [config folder](./deploy/config/).

---

### Tanka

<img src="./assets/tanka.svg" width="100px" height="100px" align="left" style="padding-right:15px"/>

Inspired by Jsonnet language, Grafana Tanka is a configuration utility that helps create kube manifests. Similar to ksonnet (now deprecated), Tanka maintains the kubernetes jsonnet library helping create DRY manifests and allows the definition of Kubernetes resources to be more concise than YAML.

| Open Source                                | Maintained By | First Release |
| ------------------------------------------ | ------------- | ------------- |
| [Repo](https://github.com/grafana/tanka) ✅ | Grafana       | 31 Jul 2019   |

```bash
make manifests.tanka
```

The command uses [jsonnet library](./deploy/tanka/lib/todo/) and [envrionments](./deploy/tanka/environments/) to template files and outputs them in [tanka folder](./deploy/output/tanka). Similar to helm, default values can be defined the app lib, while per env overrides can be defined in the environment dir using jsonnet language.

---

### ko

<img src="./assets/ko.png" width="160px" height="80px" align="left" style="padding-right:15px"/>

ko commands such as `resolve` and `apply` can hydrate your YAML manifests as ko replaces your "image:" references in YAML automatically with the image it builds.

| Open Source                            | Maintained By | First Release |
| -------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/google/ko) ✅ | Google        | 22 Mar 2019   |

```bash
# Install and template manifests using ko
make manifests.ko
```

`ko` consumes [fully baked manifests](./deploy/ko/deploy.yaml) and allows you to dynamically replace the deployment image with the exact version of image build by ko. It does this by using the following template string: `ko://01-stability-reliability/cmd/todo` where `01-stability-reliability/cmd/todo` is a go cmd path.

---

## Devex

### Skaffold

<img src="./assets/skaffold.png" width="100px" height="50px" align="left" style="padding-right:15px"/>

Skaffold handles the workflow for building, pushing and deploying your application, allowing devs to get their apps up and running quickly in a kube cluster. With the hot reload capability, it enables devs to focus on iterating on your application locally while Skaffold continuously deploys to your local or remote Kubernetes cluster. I highly advice against hot reloading and deploying apps to production cluster, had to put it in here just in case 🤯.

| Open Source                                                | Maintained By | First Release |
| ---------------------------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/GoogleContainerTools/skaffold) ✅ | Google        | 6 Mar 2018    |

```bash
make dev
```

Skaffold consumes a [skaffold.yaml](./skaffold.yaml) to define image building, tagging and deployment steps, allowing app devs to get the application up and running in a kubernetes cluster by cloning the repo and executing `skaffold dev`. Assuming you have a kube cluster running, `make dev` will:

- Install skaffold
- Create 3 namespaces
- Helm template manifests
- Create a psql DB
- Build a docker image
- Deploy the image
- Port forward to the container
- Hot reload on file changes

---

### Octant

<img src="./assets/octant.png" width="130px" height="40px" align="left" style="padding-right:15px"/>

A UI for developers which helps developers understand the Kubernetes resources deployed in a cluster. Allows for easy navigation of cluster and resources management. The pluggable nature of the UI makes it ideal for writing devex extensions.

| Open Source                                      | Maintained By | First Release |
| ------------------------------------------------ | ------------- | ------------- |
| [Repo](https://github.com/vmware-tanzu/octant) ✅ | VMware        | 20 Nov 2018   |

```bash
# Install and open octant UI
make ui
```

`make ui` runs octant which discovers kube resources installed in your kube cluster and outputs them in a localhost website, usually `http://127.0.0.1:7777/`.

---

## Deploy

### ArgoCD

<img src="./assets/argo.png" width="100px" height="100px" align="left" style="padding-right:15px"/>

Bunch of buzz words coming through:

Declarative, GitOps continuous delivery tool for Kubernetes.

Enables gitops for Kube manifests, which translates to using Git repositories as the source of truth for defining the desired application state. To specify multiple environments for a service, `ApplicationSet` controller must also be installed.

| Open Source                                   | Maintained By | First Release |
| --------------------------------------------- | ------------- | ------------- |
| [Repo](https://github.com/argoproj/argo-cd) ✅ | CNCF          | 13 Mar 2018   |

```bash
make argo.setup
make argo
```

Getting argo up and running to split into 2 commands:

- `make argo.setup`: Creates `argocd` namespaces and deploys ArgoCD and ApplicationSet controllers. ArgoCD by itself is only capable of managing 1 application in 1 environment. ApplicationSet is a controller which consumes an `ApplicationSet` CRD which enables devs to define manifest for an application deployed into multiple env / namespaces.
- You'll need to manually setup `argocd` CLI before executing the next step. You can find more information [here](https://argoproj.github.io/argo-cd/getting_started/#4-login-using-the-cli).
- `make argo`: Deploy the `ApplicationSet` [CRD](./deploy/app-set.yaml) and syncs the state of the cluster with the config defined in git. It syncs the state of resources with the files defined in [manifests folder](./deploy/manifests/)

---
