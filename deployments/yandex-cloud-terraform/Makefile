init:
	terraform init

validate:
	terraform validate

plan:
	terraform plan -var-file=terraform.tfvars -out=terraform.tfplan

apply:
	terraform apply terraform.tfplan

destroy:
	terraform destroy -var-file=terraform.tfvars

.PHONY: init validate plan apply destroy