# Recurso IAM Role para a função Lambda
resource "aws_iam_role" "lambda" {
    name = "lambda-ecr-role"

    assume_role_policy = jsonencode({
        Version   = "2012-10-17",
        Statement = [{
            Effect    = "Allow",
            Principal = {
            Service = "lambda.amazonaws.com"
            },
            Action    = "sts:AssumeRole"
        }]
    })
}

resource "aws_iam_policy" "policy_lambda" {
    name        = "loggin-${var.LAMBDA_AUTH_NAME}"
    path        = "/"
    description = "IAM policy for logging from a lambda"
    policy = jsonencode({
                    Version = "2012-10-17",
                    Statement = [{
                        Effect   = "Allow",
                        Action   = [
                            "logs:CreateLogGroup",
                            "logs:CreateLogStream",
                            "logs:PutLogEvents"
                        ],
                        Resource = "*"
                    }, 
                    {
                        Effect   = "Allow",
                        Action   = [
                            "ecr:GetAuthorizationToken",
                            "ecr:BatchCheckLayerAvailability",
                            "ecr:GetDownloadUrlForLayer",
                            "ecr:GetRepositoryPolicy",
                            "ecr:DescribeRepositories",
                            "ecr:ListImages",
                            "ecr:DescribeImages",
                            "ecr:BatchGetImage",
                            "ecr:InitiateLayerUpload",
                            "ecr:UploadLayerPart",
                            "ecr:CompleteLayerUpload",
                            "ecr:PutImage"
                        ],
                        Resource = "*"
                    }]
                })
}


resource "aws_iam_role_policy_attachment" "attachment_policy" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.policy_lambda.arn
}


# Recurso Lambda Function
resource "aws_lambda_function" "build_and_push_to_ecr" {
        function_name = var.LAMBDA_AUTH_NAME
        image_uri     = var.IMAGE_URL
        package_type  = "Image"
        timeout       = 300
        role          = aws_iam_role.lambda.arn
        publish       = true
        memory_size   = 256
        architectures = ["x86_64"]

        image_config {
            entry_point = [
                "/lambda-entrypoint.sh"
            ]
        }

        tracing_config {
            mode = "PassThrough"
        }

        environment {
            variables = {
                # ENVIRONMENT         = var.ENVIRONMENT
                # POSTGRES_HOST       = var.POSTGRES_HOST
                # POSTGRES_USER       = var.POSTGRES_USER
                # POSTGRES_PASSWORD   = var.POSTGRES_PASSWORD
                # POSTGRES_PORT       = var.POSTGRES_PORT
                # POSTGRES_DATABASE   = var.POSTGRES_DATABASE
            }
        }
}