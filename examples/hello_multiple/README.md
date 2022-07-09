# Hello Multiple
This example demoinstrates how to use the [greet_multiple](../../modules/greet_multiple/) module that extends the [Hello World!](../../) module to use a list of strings.

## Usage
1. From inside this directory run `terraform init`
2. Then run `terraform apply`

<!-- BEGIN_TF_DOCS -->
## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_greet_multiple"></a> [greet\_multiple](#module\_greet\_multiple) | ../../modules/greet_multiple | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_friends"></a> [friends](#input\_friends) | A list of people/things to greet. | `list(string)` | <pre>[<br>  "Terraform",<br>  "Terratest",<br>  "Terraform Weekly"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_result"></a> [result](#output\_result) | The resul of calling the hello\_world submodule greet\_multiple. |
<!-- END_TF_DOCS -->