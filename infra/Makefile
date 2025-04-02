init:
	terraform init -backend=false

upgrade:
	terraform init -upgrade -backend=false

validate: init
	terraform validate

plan: terraform-backend
	terraform plan

apply: terraform-backend
	terraform apply \
	-auto-approve

destroy: terraform-backend
	terraform destroy

terraform-backend:
	terraform init -reconfigure \
	  -backend=true
