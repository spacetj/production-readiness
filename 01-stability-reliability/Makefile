PROJECT_ID := sandbox-project-tc
APP_NAME := todo
CLUSTER_NAME := ${APP_NAME}
GCR_IMAGE := "gcr.io/${PROJECT_ID}/${APP_NAME}"

dev: install.skaffold env manifests.helm
	kubectl apply -f deploy/output/psql.yaml
	skaffold dev --no-prune=false --cache-artifacts=false --port-forward

build.go: install.golangci-lint
	go clean ./...
	go mod tidy && go mod vendor
	go fmt ./...
	golangci-lint run
	go build -o ./dist/todo ./cmd/todo

build.docker:
	docker build . -t ${APP_NAME}:docker

build.kaniko:
	gcloud auth application-default login
	docker run \
		-v ${HOME}/.config/gcloud:/root/.config/gcloud \
		-v $(shell PWD):/workspace \
		gcr.io/kaniko-project/executor:latest \
		--dockerfile /workspace/Dockerfile \
		--destination "${GCR_IMAGE}/kaniko" \
		--context dir:///workspace/ \
		--cache=true \
		--cache-repo "${GCR_IMAGE}/cache"

build.pack: install.pack
	pack config default-builder gcr.io/buildpacks/builder:v1
	pack build todo:pack

build.ko: install.ko
	KO_DOCKER_REPO=${GCR_IMAGE}/ko ko publish ./cmd/todo

manifests:
	helm template --name todo ./deploy/helm -f ./deploy/config/dev.yaml > ./deploy/manifests/dev/output.yaml
	helm template --name todo ./deploy/helm -f ./deploy/config/qa.yaml > ./deploy/manifests/qa/output.yaml
	helm template --name todo ./deploy/helm -f ./deploy/config/stage.yaml > ./deploy/manifests/stage/output.yaml

manifests.helm: install.helm
	rm -rf ./deploy/output/helm/*
	helm template --name todo ./deploy/helm -f ./deploy/config/dev.yaml > ./deploy/output/helm.yaml

manifests.ko: install.ko
	KO_DOCKER_REPO=${GCR_IMAGE}/ko ko resolve -f ./deploy/ko > ./deploy/output/ko.yaml

manifests.tanka: install.tanka
	rm -rf ./deploy/output/tanka/*
	tk export ./deploy/output/tanka ./deploy/tanka/environments/dev

env:
	kubectl create ns dev --dry-run -oyaml | kubectl apply -f -
	kubectl create ns stage --dry-run -oyaml | kubectl apply -f -
	kubectl create ns qa --dry-run -oyaml | kubectl apply -f -
	kubectl create ns cd --dry-run -oyaml | kubectl apply -f -

cluster: install.kind
	kind create cluster --name ${CLUSTER_NAME}

clean:
	kubectl delete -f ./output/ --ignore-not-found=true

argo: deploy.argocd argo.setup argo.sync

deploy.argocd: install.argocd
	kubectl create namespace argocd --dry-run -o yaml | kubectl apply -f -
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj-labs/applicationset/v0.1.0/manifests/install.yaml

argo.setup:
	kubectl apply -f ./deploy/app-set.yaml

argo.sync:
	argocd app sync dev
	argocd app sync qa
	argocd app sync stage

deps: install.golangci-lint install.kind install.skaffold install.ko install.tanka install.helm install.pack install.argocd

install.kind:
ifeq (, $(shell which kind))
	@echo "🤖  kind not found. installing"
	brew install kind
else 
	@echo "✅ kind found"
endif

install.golangci-lint:
ifeq (, $(shell which golangci-lint))
	@echo "🤖  golangci-lint not found. installing"
	brew install golangci-lint
else
	@echo "✅ golangci-lint found"
endif

install.skaffold:
ifeq (, $(shell which skaffold))
	@echo "🤖  skaffold not found. installing"
	brew install skaffold
else
	@echo "✅ skaffold found"
endif

install.ko:
ifeq (, $(shell which ko))
	@echo "🤖  ko not found. installing"
	brew install ko
else
	@echo "✅ ko found"
endif

install.pack:
ifeq (, $(shell which pack))
	@echo "🤖  pack not found. installing"
	brew install buildpacks/tap/pack
else
	@echo "✅ pack found"
endif

install.helm:
ifeq (, $(shell which helm))
	@echo "🤖  helm not found. installing"
	brew install helm
else
	@echo "✅ helm found"
endif

install.tanka:
ifeq (, $(shell which tk))
	@echo "🤖  tanka not found. installing"
	brew install tanka
else
	@echo "✅ tanka found"
endif
ifeq (, $(shell which jb))
	@echo "🤖  jsonnet-bundler not found. installing"
	brew install jsonnet-bundler
else
	@echo "✅ jsonnet-bundler found"
endif

install.argocd:
ifeq (, $(shell which argocd))
	@echo "🤖  argocd not found. installing"
	brew install argocd
else
	@echo "✅ argocd found"
endif
