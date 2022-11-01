variable "region" {
  description = "The AWS region"
  default     = "eu-central-1"
}

variable "availability_zones" {
  description = "A list of Availability Zones where subnets and DB instances can be created"
}

variable "deployment_stage" {
  description = "The deployment stage"
  default     = "mainnet"
}

variable "forbidden_account_ids" {
  description = "The forbidden account IDs"
  type        = list(string)
  default     = []
}

# -----------------------------------------------------------------------------
# Module service-api
# -----------------------------------------------------------------------------

variable "api_jwt_secret_ssm_key" {
  description = "Systems Manager key where the API JWT secret is stored"
}

variable "rds_satimoto_db_password_ssm_key" {
  description = "Systems Manager key where the password for the satimoto DB is stored"
}

variable "service_name" {
  description = "The name of the service"
}

variable "service_desired_count" {
  description = "The number of instances of the task definition to place and keep running"
}

variable "service_container_port" {
  description = "The port on the container to associate with the load balancer"
}

variable "service_metric_port" {
  description = "The port to associate with metric collection"
}

variable "task_network_mode" {
  description = "The Docker networking mode to use for the containers in the task"
}

variable "task_cpu" {
  description = "The number of cpu units used by the task"
}

variable "task_memory" {
  description = "The amount (in MiB) of memory used by the task"
}

variable "target_health_path" {
  description = "The path to check the target's health"
}

variable "target_health_interval" {
  description = "The approximate amount of time, in seconds, between health checks of an individual target"
}

variable "target_health_timeout" {
  description = "The amount of time, in seconds, during which no response means a failed health check"
}

variable "target_health_matcher" {
  description = "The HTTP codes to use when checking for a successful response from a target"
}

variable "subdomain_name" {
  description = "The subdomain name of the service"
}

variable "env_channel_request_max_amount" {
  description = "The maximum channel size in satoshis for a channel request"
}

variable "env_default_commission_percent" {
  description = "The environment variable to set the default commission percent"
}

variable "env_ferp_rpc_port" {
  description = "The environment variable to set the FERP RPC port"
}

variable "env_ocpi_rpc_port" {
  description = "The environment variable to set the OCPI RPC port"
}

variable "env_shutdown_timeout" {
  description = "The environment variable to set the shutdown timeout"
}
