.PHONY: build-image build-helm
.ONESHELL:

export TAG ?= $(shell git rev-parse --short HEAD)

build-images:
	@eval $$(minikube docker-env) ;\
	docker build -t orchestrator:$(TAG) ../src
	
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
	cd orchestrator && helm dependency build

start-mk:
	minikube start \
		--cpus=4 \
		--memory=4096 \
		--disk-size=20000mb \
		--kubernetes-version 1.18.3 \
		--vm-driver virtualbox ;

start: start-mk install-deps deploy-minikube

restart:
	make cleanup
	make start

cleanup-minikube:
	minikube stop
	minikube delete

kashti:
	minikube service minikube-kashti

fireworq:
	kubectl port-forward service/fireworq 8080:8080

generic-gateway:
	kubectl port-forward service/minikube-brigade-generic-gateway 8081:8081