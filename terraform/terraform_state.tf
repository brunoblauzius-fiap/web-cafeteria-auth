terraform {
  backend "s3" {
    bucket = "terraform-clickmark-tf"
    key    = "terraform-tech-infra/terraform.tfstate"
    region = "us-east-1"
  }
}