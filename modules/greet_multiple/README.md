# Greet Multiple
Greet multiple is a Higher-order module, which is Terraform's version of a [Higher-Order Function](https://en.wikipedia.org/wiki/Higher-order_function). Greet Multiple extends the capability of the [Hello World!](../../) module by changing the input from a single string to a list of strings.

While you could modify [Hello World!](../../) to accept both a single string and a list of strings, doing so would cause a significant increase in complexity. Making future changes harder to implement and tests harder to write.

A real-world use case for Higher-order modules would be if you wanted your module to create a resource or allow the resource id to be passed in. Have your "main" module accept the resource id as a variable. Then create a submodule to create the resource and call the "main" module with the resource id that it created.

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