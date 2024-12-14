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

resource "aws_ecr_repository" "repository" {
  name                 = "go-ecr-deploy"
  image_tag_mutability = "MUTABLE"
}

resource "aws_iam_user" "ecr_user" {
  name = "go-ecr-deploy"
}

resource "aws_iam_policy" "ecr_publish_policy" {
  name        = "go-ecr-deploy-publish-policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ecr:BatchCheckLayerAvailability",
          "ecr:PutImage",
          "ecr:InitiateLayerUpload",
          "ecr:UploadLayerPart"
        ]
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_user_policy_attachment" "ecr_user_policy" {
  user       = aws_iam_user.ecr_user.name
  policy_arn = aws_iam_policy.ecr_publish_policy.arn
}

resource "aws_iam_access_key" "ecr_user_access_key" {
  user = aws_iam_user.ecr_user.name
}

output "AWS_ACCESS_KEY_ID" {
  value = aws_iam_access_key.ecr_user_access_key.id
}

output "AWS_SECRET_ACCESS_KEY" {
  value = aws_iam_access_key.ecr_user_access_key.secret
  sensitive = true
}