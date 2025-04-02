data "archive_file" "flush_cache" {
  type        = "zip"
  source_file = "src/lambda.py"
  output_path = "src/lambda_function_payload.zip"
}

resource "aws_lambda_function" "flush_cache" {
  filename      = "lambda_function_payload.zip"
  function_name = var.lambda_name
  role          = aws_iam_role.flush_cache.arn
  handler       = "lambda.lambda_handler"

  runtime          = var.lambda_runtime
  source_code_hash = data.archive_file.flush_cache.output_base64sha256

}

resource "aws_iam_role" "flush_cache" {
  name               = "${var.lambda_name}-lambda"
  description        = "IAM role for ${var.lambda_name} lambda" 
  assume_role_policy = data.aws_iam_policy_document.flush_cache_assume_role.json
}

data "aws_iam_policy_document" "flush_cache_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role_policy_attachment" "flush_cache" {
  role       = aws_iam_role.flush_cache.name
  policy_arn = aws_iam_policy.flush_cache.arn
}

resource "aws_iam_policy" "flush_cache" {
  name        = "${var.lambda_name}-lambda"
  path        = "/"
  description = "IAM policy for ${var.lambda_name} lambda"
  policy      = data.aws_iam_policy_document.flush_cache_policy.json
}

data "aws_iam_policy_document" "flush_cache_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
    ]
    resources = ["arn:aws:logs:${data.aws_region.this.name}:${data.aws_caller_identity.this.account_id}:*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resources = ["arn:aws:logs:${data.aws_region.this.name}:${data.aws_caller_identity.this.account_id}:log-group:/aws/lambda/${var.lambda_name}:*"]
  }
}
