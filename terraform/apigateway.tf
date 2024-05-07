resource "aws_api_gateway_rest_api" "prod" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "fiap api - prod"
      version = "1.0"
    }
    paths = {
      "/auth" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "POST"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "fiap-api-prod"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_deployment" "prod" {
  rest_api_id = aws_api_gateway_rest_api.prod.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.prod.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "prod" {
  deployment_id = aws_api_gateway_deployment.prod.id
  rest_api_id   = aws_api_gateway_rest_api.prod.id
  stage_name    = "prod"
}