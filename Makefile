# this Makefile depends on
# scripts from https://github.com/jcosentino11/scripts

# run the stack locally
run-local:
	@aws_login local
	docker-compose up

# login to aws-cli
login-staging: logout
	@aws_login meals-deploy-staging

login-production: logout
	@aws_login meals-deploy-production

# logout of aws-cli
logout:
	@aws_logout

# apply all the cloudbyte infra
apply-staging: login-staging
	@with_terraform "cd infra/staging && terraform init && terraform get && terraform apply"

# destroy all the cloudbyte infra
destroy-staging: login-staging
	@with_terraform "cd infra/staging && terraform init && terraform get && terraform destroy"

