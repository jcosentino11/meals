variable "user_pool_allow_signups" {
  type        = bool
  default     = false
  description = "If true, users can sign themselves up."
}

variable "user_pool_client_refresh_token_expiration_days" {
  type        = number
  default     = 30
  description = "Refresh token expiration, in days."
}

variable "application_name" {
  type        = string
  description = "The name of the application."
}

variable "environment" {
  type        = string
  description = "Name of the application environment (e.g. staging, production)."
}
