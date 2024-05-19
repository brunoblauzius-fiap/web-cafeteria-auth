provider "aws" {
  region = var.REGION
}

resource "aws_lambda_function" "lambda_function" {
  filename         = "${path.module}/deploy.zip"
  function_name    = "authentication"
  role             = aws_iam_role.lambda_role.arn
  handler          = "bootstrap"
  source_code_hash = filebase64sha256("${path.module}/deploy.zip")
  runtime          = "go1.x"
}

resource "aws_lambda_function" "gera_token" {
  function_name    = "gera-token"
  role             = aws_iam_role.lambda_role.arn
  handler          = "lambda_function.lambda_handler"
  runtime          = "python3.8"
  filename         = "${path.module}/../lambda/lambda_function.zip"
  source_code_hash = filebase64sha256("${path.module}/../lambda/lambda_function.zip")
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_execution_role"
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Principal = {
        Service = "lambda.amazonaws.com"
      },
      Effect = "Allow",
      Sid    = ""
    }]
  })

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecu
