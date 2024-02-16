build:
	@docker build -t local_kind_app .
run:
	@docker run -p 8080:8080 -e PODNAME=pod-123 -e PODNAMESPACE=default -e NODENAME=kind-development local_kind_app
