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
  name = "go-ecr-deploy-publish-policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      // ref: https://github.com/aws-actions/amazon-ecr-login?tab=readme-ov-file#permissions
      {
        Effect = "Allow",
        Action = [
          "ecr:GetAuthorizationToken"
        ],
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ecr:BatchGetImage",
          "ecr:BatchCheckLayerAvailability",
          "ecr:CompleteLayerUpload",
          "ecr:GetDownloadUrlForLayer",
          "ecr:InitiateLayerUpload",
          "ecr:PutImage",
          "ecr:UploadLayerPart"
        ]
        Resource = aws_ecr_repository.repository.arn
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

resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Effect = "Allow"
        Sid    = ""
      }
    ]
  })
}

resource "aws_lambda_function" "lambda" {
  function_name = "go-ecr-deploy-example"
  role          = aws_iam_role.lambda_exec_role.arn
  package_type  = "Image"
  image_uri    = "${aws_ecr_repository.repository.repository_url}:multistage-copy-nocache"
}

resource "aws_lambda_function_url" "lambda_function_url" {
  function_name = aws_lambda_function.lambda.function_name
  authorization_type = "NONE"
}

output "lambda_function_url" {
  value = aws_lambda_function_url.lambda_function_url.function_url
}

output "AWS_ACCESS_KEY_ID" {
  value = aws_iam_access_key.ecr_user_access_key.id
}

output "AWS_SECRET_ACCESS_KEY" {
  value     = aws_iam_access_key.ecr_user_access_key.secret
  sensitive = true
}