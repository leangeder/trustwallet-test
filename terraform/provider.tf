terraform {
  required_version = ">= 1.9.2"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.58"
    }
  }

  backend "local" {}
}

data "aws_availability_zones" "available" {}

provider "aws" {
  region = local.region
}
