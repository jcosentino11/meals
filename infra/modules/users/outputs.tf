output "values" {
  value = {
    user_pool_id = aws_cognito_user_pool.pool.id
    user_pool_app_client_id = aws_cognito_user_pool_client.client.id
    user_pool_domain_id = aws_cognito_user_pool_domain.domain.id
  }
}
