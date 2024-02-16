.PHONY: deploy
deploy:
	@kubectl apply -f manifests/deployment.yaml

.PHONY: delete
delete:
	@kubectl delete -f manifests/deployment.yaml

.PHONY: load_image
load_image:
	@kind load docker-image local_kind_app:latest --name development

.PHONY: get_logs
get_logs:
	@kubectl logs -l app=app

.PHONY: request
request:
	@curl http://localhost:8080

.PHONY: port_forward
port_forward:
	pod_name=$$(kubectl get pods --no-headers -o custom-columns=":metadata.name" | head -n 1); \
    kubectl port-forward $$pod_name 8080:8080