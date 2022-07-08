# Hello Custom!
This example demoinstrates how to run the module using custom input values and shows what output is to be expected.

## Usage
1. From inside this directory run `terraform init`
2. Then run `terraform apply`
3. When prompted for the `name` value, enter your name.

<!-- BEGIN_TF_DOCS -->
## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_hello_world"></a> [hello\_world](#module\_hello\_world) | ../../ | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Enter your name. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_result"></a> [result](#output\_result) | The resul of calling the hello\_world module by passing the name arg. |
<!-- END_TF_DOCS -->