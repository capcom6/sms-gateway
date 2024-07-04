SHELL:=/bin/bash
TF_DIR=.
TF_PLAN=${TF_DIR}/terraform.tfplan

.PHONY: init plan apply destroy

init:
	cd ${TF_DIR} && terraform init

plan:
	cd ${TF_DIR} && terraform plan -var-file="secrets.tfvars" -out=${TF_PLAN}

apply:
	cd ${TF_DIR} && terraform apply ${TF_PLAN}

destroy:
	cd ${TF_DIR} && terraform destroy -var-file="secrets.tfvars"
