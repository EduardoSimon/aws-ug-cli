terraform {
  required_version = "~> 1.0"
  backend "local" {}
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  default_tags {
    tags = {
      app     = "aws-ug-cli"
      managed = "terraform"
      url     = "https://github.com/EduardoSimon/aws-ug-cli"
    }
  }
}
