region = "eu-central-1"

availability_zones = ["eu-central-1a", "eu-central-1b", "eu-central-1c"]

deployment_stage = "testnet"

forbidden_account_ids = ["490833747373"]

# -----------------------------------------------------------------------------
# Module service-api
# -----------------------------------------------------------------------------

api_jwt_secret_ssm_key = "/api/jwt_secret"

rds_satimoto_db_password_ssm_key = "/rds/satimoto_db_password"

service_name = "api"

service_desired_count = 1

service_container_port = 9000

service_metric_port = 9100

task_network_mode = "awsvpc"

task_cpu = 256

task_memory = 512

target_health_path = "/health"

target_health_interval = 120

target_health_timeout = 5

target_health_matcher = "200"

subdomain_name = "api"

env_channel_request_max_amount = 450000

env_default_commission_percent = 7

env_fcm_api_key = "AAAAxi-5AuI:APA91bGGRy1tshxXQuqyz3vjbarr5QFwlGpKJAUfoAQ-IqCU_8kWClNslZpRKEExQtew2YxQwtIyOIe9QDCCbbApPoqAO9Srm2i-CNNM8agj2mkMJsLMaEYGtGW81QYpvQJqwBSQWVYf"

env_ferp_rpc_port = 50000

env_ocpi_rpc_port = 50000

env_shutdown_timeout = 20
