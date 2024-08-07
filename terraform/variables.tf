variable "name" {
  type    = string
  default = "test"
}

locals {
  suffix = "trustwallet"
  name   = join("-", [var.name, local.suffix])
  region = "eu-west-1"

  vpc_cidr = "10.0.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)

  container_name = "ecsdemo-frontend"
  container_port = 9090
}
