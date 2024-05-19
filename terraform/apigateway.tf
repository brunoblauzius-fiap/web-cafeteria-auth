resource "aws_api_gateway_rest_api" "prod" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "fiap api - prod"
      version = "1.0"
    }
    # paths = {
    #   "/auth" = {
    #     post = {
    #       x-amazon-apigateway-integration = {
    #         httpMethod           = "POST"
    #         type                 = "AWS_PROXY"
    #         uri                  = "${aws_lambda_function.gera_token.invoke_arn}"
    #       }
    #     }
    #   }
    # }
  })

  name = "fiap-api-prod"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}


resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.gera_token.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.prod.execution_arn}/*/*"
}

resource "aws_api_gateway_resource" "auth" {
  rest_api_id = aws_api_gateway_rest_api.prod.id
  parent_id   = aws_api_gateway_rest_api.prod.root_resource_id
  path_part   = "auth"
}

resource "aws_api_gateway_method" "auth_post" {
  rest_api_id   = aws_api_gateway_rest_api.prod.id
  resource_id   = aws_api_gateway_resource.auth.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "auth_post" {
  rest_api_id = aws_api_gateway_rest_api.prod.id
  resource_id = aws_api_gateway_resource.auth.id
  http_method = aws_api_gateway_method.auth_post.http_method
  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = aws_lambda_function.gera_token.invoke_arn
}

resource "aws_api_gateway_deployment" "prod" {
  depends_on = [aws_api_gateway_integration.auth_post]
  rest_api_id = aws_api_gateway_rest_api.prod.id
  stage_name  = "prod"
}