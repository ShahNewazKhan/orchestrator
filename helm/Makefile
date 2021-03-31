.PHONY: build-image build-helm
.ONESHELL:

export TAG ?= $(shell git rev-parse --short HEAD)

build-images:
	@eval $$(minikube docker-env) ;\
	echo $$CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
	DOCKER_BUILDKIT=1 docker build -t orchestrator:$(TAG) --network=host ../src
	DOCKER_BUILDKIT=1 docker build -t coeus-engine:$(TAG) -t coeus-engine:latest ../../coeus-engine
	
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

create-aws-secrets:
	if [ -z AWS_ACCESS_KEY_ID ]; then echo "AWS_ACCESS_KEY_ID is unset" && exit 1; else echo "AWS_ACCESS_KEY_ID is set to '$$AWS_ACCESS_KEY_ID'"; fi
	if [ -z AWS_SECRET_ACCESS_KEY ]; then echo "AWS_SECRET_ACCESS_KEY is unset" && exit 1; else echo "AWS_SECRET_ACCESS_KEY is set to '$$AWS_SECRET_ACCESS_KEY'"; fi
	kubectl create secret generic aws-access-key-id --from-literal=access-key-id=$$AWS_ACCESS_KEY_ID
	kubectl create secret generic aws-secret-access-key --from-literal=secret-access-key=$$AWS_SECRET_ACCESS_KEY

get-aws-creds:
	kubectl get secrets/aws-access-key-id --template={{.data.access-key-id}}

start: start-mk install-deps deploy-minikube

restart:
	make cleanup-minikube
	make start

cleanup-minikube:
	minikube stop
	minikube delete

kashti:
	minikube service minikube-kashti

kasti-fwd:
	kubectl port-forward service/minikube-kashti 8081:80

generic-gateway:
	kubectl port-forward service/minikube-brigade-generic-gateway 8081:8081

brigade-api:
	kubectl port-forward service/minikube-brigade-api 7745:7745

local-dev: kashti brigade-api

connect-kubesail:
	kubectl create -f https://byoc.kubesail.com/shahnewazkhan.yaml

generic-gw-secret:
	kubectl get secrets/brigade-6ed38d031ae3403c0608d0d5b2361a0e4e27ef2dedf2810291c433 --template={{.data.genericGatewaySecret}} | base64
