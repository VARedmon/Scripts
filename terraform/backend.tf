terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}

# For team environments, replace the local backend with a remote backend.
# Example:
# terraform {
#   backend "s3" {
#     bucket = "my-terraform-state"
#     key    = "envs/prod/terraform.tfstate"
#     region = "us-east-1"
#   }
# }
