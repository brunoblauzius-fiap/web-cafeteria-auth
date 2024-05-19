provider "aws" {
  region = var.REGION
}

resource "aws_lambda_function" "gera_token" {
  function_name = "gera-token"
  role          = aws_iam_role.lambda_role.arn
  # handler       = "lambda_function.lambda_handler"
  # runtime       = "python3.8"
  handler       = "main"  # O handler é o nome do executável Go
  runtime       = "provided.al2"

  filename      = "${path.module}/../authentication/lambda_function.zip"
  source_code_hash = filebase64sha256("${path.module}/../authentication/lambda_function.zip")
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect    = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Action    = "sts:AssumeRole"
      }
    ]
  })

#   policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Effect": "Allow",
#       "Action": [
#         "logs:CreateLogGroup",
#         "logs:CreateLogStream",
#         "logs:PutLogEvents"
#       ],
#       "Resource": "arn:aws:logs:*:*:*"
#     }
#   ]
# }
# EOF
}