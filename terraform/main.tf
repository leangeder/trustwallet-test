module "ecs" {
  source = "terraform-aws-modules/ecs/aws"

  cluster_name = local.name

  # Capacity provider
  fargate_capacity_providers = {
    FARGATE = {
      default_capacity_provider_strategy = {
        weight = 50
        base   = 20
      }
    }
    FARGATE_SPOT = {
      default_capacity_provider_strategy = {
        weight = 50
      }
    }
  }

  services = {
    trustwallet_app = {
      cpu    = 512
      memory = 1024

      # Container definition(s)
      container_definitions = {

        (local.container_name) = {
          cpu       = 512
          memory    = 512
          essential = true
          image     = "public.ecr.aws/trustwallet/app:latest"

          health_check = {
            command = ["CMD-SHELL", "curl -f http://localhost:${local.container_port}/healthz || exit 1"]
          }

          port_mappings = [
            {
              name          = local.container_name
              containerPort = local.container_port
              hostPort      = local.container_port
              protocol      = "tcp"
            }
          ]

          # Example image used requires access to write to root filesystem
          readonly_root_filesystem = true

          enable_cloudwatch_logging = false
          memory_reservation        = 100
        }
      }

      subnet_ids = module.vpc.private_subnets
    }
  }
}
