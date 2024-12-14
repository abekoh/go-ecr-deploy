terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.81"
    }
  }

  required_version = ">= 1.9.7"
}

provider "aws" {
  region = "us-west-2"
}

resource "aws_ecr_repository" "foo" {
  name                 = "go-ecr-deploy"
  image_tag_mutability = "MUTABLE"
}
