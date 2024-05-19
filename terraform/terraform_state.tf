terraform {
  backend "s3" {
    bucket = "terraform-clickmark-tf"
    key    = "terraform-webauth/terraform.tfstate"
    region = "us-east-1"
  }
}