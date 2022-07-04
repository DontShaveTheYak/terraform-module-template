# Terraform module template

This "Hello World" module is a template and a guide to creating Terraform modules.

### Features
* Local testing with pre-commit
    - Terraform Format
    - Terraform Validate
    - Automatic `README` Updates with `terraform-docs`
    - Static code analysis with `tflint`
    - Static code analysis with `tfsec`
    - Static code analysis with `checkov`
* Devcontainer with depencies pre-configured
* CICD pipeline with Github Actions
* Testing with `Terratest` across all major versions of Terraform.
* Shows use of module composition

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.12.0, < 2.0.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_make_exciting"></a> [make\_exciting](#module\_make\_exciting) | ./modules/make_exciting | n/a |
| <a name="module_say_hello"></a> [say\_hello](#module\_say\_hello) | ./modules/say_hello | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | The name of a person or thing to say hello to. | `string` | `"World"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_greeting"></a> [greeting](#output\_greeting) | A very exciting greeting! |
<!-- END_TF_DOCS -->