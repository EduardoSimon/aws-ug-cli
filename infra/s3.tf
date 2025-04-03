resource "aws_s3_bucket" "demo" {
  bucket        = var.s3_name
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "demo" {
  bucket = aws_s3_bucket.demo.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_ownership_controls" "demo" {
  bucket = aws_s3_bucket.demo.id

  rule {
    object_ownership = "BucketOwnerEnforced"
  }
}

resource "aws_s3_object" "config" {
  for_each = toset([for f in fileset("data", "*.config") : trimsuffix(f, ".config")])

  bucket = aws_s3_bucket.demo.id
  key    = "apps/config/${each.key}/config"
  source = "data/${each.key}.config"
  etag   = filemd5("data/${each.key}.config")
}
