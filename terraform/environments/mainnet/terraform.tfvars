region = "eu-central-1"

availability_zones = ["eu-central-1a", "eu-central-1b", "eu-central-1c"]

deployment_stage = "mainnet"

forbidden_account_ids = ["909899099608"]

# -----------------------------------------------------------------------------
# Module service-api
# -----------------------------------------------------------------------------

api_jwt_secret_ssm_key = "/api/jwt_secret"

rds_satimoto_db_password_ssm_key = "/rds/satimoto_db_password"

service_name = "api"

service_desired_count = 1

service_container_port = 9000

task_network_mode = "awsvpc"

task_cpu = 256

task_memory = 512

target_health_path = "/health"

target_health_interval = 120

target_health_timeout = 5

target_health_matcher = "200"

subdomain_name = "api"

env_default_commission_percent = 7

env_ferp_rpc_port = 50000

env_ocpi_rpc_port = 50000

env_shutdown_timeout = 20
