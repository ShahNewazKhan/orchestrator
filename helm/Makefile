.PHONY: build-image build-helm
.ONESHELL:

export TAG ?= $(shell git rev-parse --short HEAD)

build-images:
	@eval $$(minikube docker-env) ;\
	echo $$CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
	docker build -t orchestrator:$(TAG) --network=host ../src
	DOCKER_BUILDKIT=1 docker build -t coeus-engine:$(TAG) ../../coeus-engine
	
get-tag:
	echo $(TAG)

deploy-minikube: build-images
	kubectl config use-context minikube; \
	helm upgrade minikube ./orchestrator \
		--install \
		--set orchestrator.image=orchestrator:$(TAG)
		# --debug

delete-minikube:
	kubectl config use-context minikube; \
	helm uninstall minikube

install-deps: 
	helm repo add brigade https://brigadecore.github.io/charts
	helm repo add bitnami https://charts.bitnami.com/bitnami
	cd orchestrator && helm dependency update && helm dependency build

start-mk:
	minikube start \
		--cpus=4 \
		--memory=4096 \
		--disk-size=20000mb \
		--kubernetes-version 1.19.1 \
		--driver virtualbox ;

start: start-mk install-deps deploy-minikube

restart:
	make cleanup-minikube
	make start

cleanup-minikube:
	minikube stop
	minikube delete

kashti:
	minikube service minikube-kashti

generic-gateway:
	kubectl port-forward service/minikube-brigade-generic-gateway 8081:8081

brigade-api:
	kubectl port-forward service/minikube-brigade-api 7745:7745

local-dev: kashti brigade-api

connect-kubesail:
	kubectl create -f https://byoc.kubesail.com/shahnewazkhan.yaml

generic-gw-secret:
	kubectl get secrets/brigade-6ed38d031ae3403c0608d0d5b2361a0e4e27ef2dedf2810291c433 --template={{.data.genericGatewaySecret}} | base64
