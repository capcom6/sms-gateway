init:
	terraform init

validate:
	terraform validate

plan:
	terraform plan -var-file=secrets.tfvars -out=terraform.tfplan

apply:
	terraform apply terraform.tfplan

destroy:
	terraform destroy -var-file=secrets.tfvars

.PHONY: init validate plan apply destroy