provider "aws" {
  region                = var.region
  forbidden_account_ids = var.forbidden_account_ids
  profile               = "satimoto-mainnet"
}

provider "aws" {
  alias                 = "us_east_1"
  region                = "us-east-1"
  forbidden_account_ids = var.forbidden_account_ids
  profile               = "satimoto-mainnet"
}

provider "aws" {
  alias                 = "zone_owner"
  region                = var.region
  forbidden_account_ids = var.forbidden_account_ids
  profile               = "satimoto-common"
}

# -----------------------------------------------------------------------------
# Backends
# -----------------------------------------------------------------------------

data "terraform_remote_state" "infrastructure" {
  backend = "s3"

  config = {
    bucket  = "satimoto-terraform-mainnet"
    key     = "infrastructure.tfstate"
    region  = "eu-central-1"
    profile = "satimoto-mainnet"
  }
}

terraform {
  backend "s3" {
    bucket  = "satimoto-terraform-mainnet"
    key     = "api.tfstate"
    region  = "eu-central-1"
    profile = "satimoto-mainnet"
  }
}

# -----------------------------------------------------------------------------
# Modules
# -----------------------------------------------------------------------------

module "subdomain_zone" {
  providers = {
    aws.zone_owner = aws.zone_owner
  }
  source             = "git::https://github.com/satimoto/terraform-infrastructure.git//modules/subdomain-zone?ref=develop"
  availability_zones = var.availability_zones
  region             = var.region

  domain_name     = data.terraform_remote_state.infrastructure.outputs.route53_zone_name
  subdomain_name  = var.subdomain_name
  route53_zone_id = data.terraform_remote_state.infrastructure.outputs.route53_zone_id
}

data "aws_caller_identity" "current" {}

data "aws_ssm_parameter" "jwt_secret" {
  name = var.api_jwt_secret_ssm_key
}

data "aws_ssm_parameter" "satimoto_db_password" {
  name = var.rds_satimoto_db_password_ssm_key
}

resource "aws_iam_role" "task_role" {
  name               = "${var.service_name}-task-role"
  assume_role_policy = file("../../resources/ecs-task-role.json")

  inline_policy {
    name = "${var.service_name}-ses-task-role-policy"
    policy = jsonencode({
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "ses:SendEmail",
            "ses:SendRawEmail"
          ],
          "Resource" : "*"
        }
      ]
    })
  }
}

module "service-api" {
  source             = "git::https://github.com/satimoto/terraform-infrastructure.git//modules/service?ref=f9cad99f17c1d7c14273b9433e249922a2b92544"
  availability_zones = var.availability_zones
  deployment_stage   = var.deployment_stage
  region             = var.region

  vpc_id                         = data.terraform_remote_state.infrastructure.outputs.vpc_id
  private_subnet_ids             = data.terraform_remote_state.infrastructure.outputs.private_subnet_ids
  route53_zone_id                = module.subdomain_zone.route53_zone_id
  alb_security_group_id          = data.terraform_remote_state.infrastructure.outputs.alb_security_group_id
  alb_dns_name                   = data.terraform_remote_state.infrastructure.outputs.alb_dns_name
  alb_zone_id                    = data.terraform_remote_state.infrastructure.outputs.alb_zone_id
  alb_listener_arn               = data.terraform_remote_state.infrastructure.outputs.alb_listener_arn
  ecs_cluster_id                 = data.terraform_remote_state.infrastructure.outputs.ecs_cluster_id
  ecs_security_group_id          = data.terraform_remote_state.infrastructure.outputs.ecs_security_group_id
  ecs_task_execution_role_arn    = data.terraform_remote_state.infrastructure.outputs.ecs_task_execution_role_arn
  service_name                   = var.service_name
  service_domain_names           = ["${var.subdomain_name}.${data.terraform_remote_state.infrastructure.outputs.route53_zone_name}"]
  service_desired_count          = var.service_desired_count
  service_container_name         = var.service_name
  service_container_port         = var.service_container_port
  task_network_mode              = var.task_network_mode
  task_cpu                       = var.task_cpu
  task_memory                    = var.task_memory
  task_role_arn                  = aws_iam_role.task_role.arn
  target_health_path             = var.target_health_path
  target_health_interval         = var.target_health_interval
  target_health_timeout          = var.target_health_timeout
  target_health_matcher          = var.target_health_matcher
  service_discovery_namespace_id = data.terraform_remote_state.infrastructure.outputs.ecs_service_discovery_namespace_id

  task_container_definitions = templatefile("../../resources/task-container-definitions.json", {
    account_id                     = data.aws_caller_identity.current.account_id
    image_tag                      = "mainnet"
    region                         = var.region
    service_name                   = var.service_name
    service_container_port         = var.service_container_port
    task_network_mode              = var.task_network_mode
    env_api_domain                 = "https://${var.subdomain_name}.${data.terraform_remote_state.infrastructure.outputs.route53_zone_name}"
    env_db_user                    = "satimoto"
    env_db_pass                    = data.aws_ssm_parameter.satimoto_db_password.value
    env_db_host                    = "${data.terraform_remote_state.infrastructure.outputs.rds_cluster_endpoint}:${data.terraform_remote_state.infrastructure.outputs.rds_cluster_port}"
    env_db_name                    = "satimoto"
    env_channel_request_max_amount = var.env_channel_request_max_amount
    env_default_commission_percent = var.env_default_commission_percent
    env_ferp_rpc_address           = "ferp.${data.terraform_remote_state.infrastructure.outputs.ecs_service_discovery_namespace_name}:${var.env_ferp_rpc_port}"
    env_jwt_secret                 = data.aws_ssm_parameter.jwt_secret.value
    env_ocpi_rpc_address           = "ocpi.${data.terraform_remote_state.infrastructure.outputs.ecs_service_discovery_namespace_name}:${var.env_ocpi_rpc_port}"
    env_reply_to_email             = "Satimoto <hello@satimoto.com>"
    env_rest_port                  = var.service_container_port
    env_shutdown_timeout           = var.env_shutdown_timeout
    env_web_domain                 = "https://${data.terraform_remote_state.infrastructure.outputs.route53_zone_name}"
  })
}
