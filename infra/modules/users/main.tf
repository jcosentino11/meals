resource "aws_cognito_user_pool" "pool" {
  name = "${var.application_name}-${var.environment}-user-pool"

  admin_create_user_config {
    allow_admin_create_user_only = ! var.user_pool_allow_signups
  }

  username_attributes = ["email"]

  password_policy {
    minimum_length = 8
    temporary_password_validity_days = 1
  }

  verification_message_template {
    default_email_option  = "CONFIRM_WITH_LINK"
    email_subject_by_link = "Verify your ${var.application_name} email address"
    email_message_by_link = "Please verify your email address. {##Click here to verify##}"
  }

  tags = {
    Name        = "${var.application_name}-${var.environment}-user-pool"
    Application = var.application_name
    Description = "User pool for ${var.application_name} application."
    Environment = var.environment
  }
}

resource "aws_cognito_user_pool_client" "client" {
  name         = "${var.application_name}-${var.environment}"
  user_pool_id = aws_cognito_user_pool.pool.id

  refresh_token_validity = var.user_pool_client_refresh_token_expiration_days

  explicit_auth_flows = [
    "ALLOW_ADMIN_USER_PASSWORD_AUTH",
    "ALLOW_CUSTOM_AUTH",
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]

}

resource "aws_cognito_user_pool_domain" "domain" {
  domain       = "${var.application_name}-${var.environment}"
  user_pool_id = aws_cognito_user_pool.pool.id
}
