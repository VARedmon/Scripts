# Terraform Project

This directory is structured for larger Terraform projects.

## Recommended layout

- `providers.tf` — provider and required version configuration
- `backend.tf` — backend configuration for remote state
- `variables.tf` — variable definitions and defaults
- `terraform.tfvars.example` — example input values
- `outputs.tf` — exported outputs
- `modules/` — reusable Terraform modules
- `environments/` — environment-specific configurations
- `.gitignore` — ignored Terraform artifacts

## Usage

1. `terraform init`
2. `terraform plan`
3. `terraform apply`

For team projects, configure a remote backend in `backend.tf`.
