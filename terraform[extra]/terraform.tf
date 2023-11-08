resource "aws_iam_role" "lambda_role" {
  name = "lambda_execution_role"
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_lambda_function" "tech-challenge-auth-lambda" {
  function_name = "tech-challenge-auth-lambda"
  package_type = "Image"
  image_uri = "564521364654.dkr.ecr.sa-east-1.amazonaws.com/tech-challenge-auth"
  
  role = aws_iam_role.lambda_role.arn
}
