terraform {
  backend "s3" {
    region   = "us-east-1"
    bucket   = "cloudbyte-tfstate-staging"
    key      = "meals/main.tfstate"
    role_arn = "arn:aws:iam::965778745441:role/meals-deploy-staging"
  }
  required_version = "0.12.24"
  required_providers {
    aws = ">= 2.58.0"
  }
}

provider "aws" {
  assume_role {
    role_arn = "arn:aws:iam::965778745441:role/meals-deploy-staging"
  }
  region = "us-east-1"
}

locals {
  application_name = "meals"
  environment      = "staging"
}

module "users" {
  source = "./../modules/users"

  application_name = local.application_name
  environment      = local.environment

  user_pool_allow_signups = false
}
