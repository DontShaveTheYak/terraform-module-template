# Greet Multiple
Greet multiple is a Higher-order module, which is Terraform's version of a [Higher-Order Function](https://en.wikipedia.org/wiki/Higher-order_function). Greet Multiple extends the capabiliy of the [Hello World!](../../) module by changing the input from a string to a list of strings.

<!-- BEGIN_TF_DOCS -->
## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_hello"></a> [hello](#module\_hello) | ../../ | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_names"></a> [names](#input\_names) | A list of people/things to greet. | `list(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_greeting"></a> [greeting](#output\_greeting) | A friendly greeting to all of our friends! |
<!-- END_TF_DOCS -->