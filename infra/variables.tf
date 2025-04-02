variable "aws_region" {
  description = "AWS region where resources will be deployed."
  type        = string
  default     = "eu-west-1"
}

variable "s3_name" {
  description = "S3 bucket name."
  type        = string
  default     = "awsugvlc-apps-config"
}

variable "db_name" {
  description = "DynamoDB table name."
  type        = string
  default     = "products"
}

variable "lambda_name" {
  description = "Lambda function name."
  type        = string
  default     = "flush-cache"
}

variable "lambda_runtime" {
  description = "Lambda function runtime."
  type        = string
  default     = "python3.13"
}
