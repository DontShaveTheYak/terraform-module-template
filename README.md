<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Python][terraform-shield]][tf-version]
[![Latest][version-shield]][release-url]
[![Tests][test-shield]][test-url]
[![License][license-shield]][license-url]
<!-- [![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url] -->

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <!-- <a href="https://github.com/DontShaveTheYak/terraform-module-template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <h3 align="center">Terraform module template</h3>

  <p align="center">
    Create a Terraform module with CICD already setup.
    <!-- <br />
    <a href="https://github.com/DontShaveTheYak/terraform-module-template"><strong>Explore the docs »</strong></a>
    <br /> -->
    <br />
    <!-- <a href="https://github.com/DontShaveTheYak/terraform-module-template">View Demo</a>
    · -->
    <a href="https://github.com/DontShaveTheYak/terraform-module-template/issues">Report Bug</a>
    ·
    <a href="https://github.com/DontShaveTheYak/terraform-module-template/issues">Request Feature</a>
    ·
    <!-- <a href="https://la-tech.co/post/hypermodern-cloudformation/getting-started/">Guide</a> -->
  </p>
</p>

## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

This "Hello World" module is a template and a guide to creating Terraform modules. Its purpose is to demonstrate Terraform module best practices and provide a template with CI/CD preconfigured.

### Features
* Local testing with [pre-commit-terraform].
    - Formats code with `terraform fmt`.
    - Validates code with `terraform validate`.
    - Automatic `README` updates with [terraform-docs].
    - Static code analysis with [TFLint], [tfsec] and [checkov].
* [Devcontainer](https://code.visualstudio.com/docs/remote/containers) with dependencies pre-configured.
* CI pipeline with Github Actions:
    * Reuses Devcontainer to run the same checks that run locally.
    * Runs tests against the latest Terraform version.
    * Runs tests against older versions of Terraform.
* CD pipeline with Github Actions:
    * Create git tags using [semver](https://semver.org/) when PRs merge to the `main` branch.
    * Create [Github Releases](https://github.com/DontShaveTheYak/terraform-module-template/releases) for every tag.
* Testing with [Terratest] across all major versions of Terraform.
* Shows use of module composition by combining smaller modules
* Shows how to use "Higher-order modules" to [extend module functionality](./modules/greet_multiple/).

## Getting Started

## Usage
If you are creating a brand new module then using this template is pretty straight forward by using the GitHub [guide](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template).

If you are wanting to apply this template to an existing module then you will want to do the following:
1. Copy the [.devcontainer](.devcontainer) and [.github](.github) directories.
2. Update your README file/files to include the [template value](https://terraform-docs.io/user-guide/configuration/output/) where you want the auto generated content from [terraform-docs] to go.

Regardless if you created a new module or updated an existing one, you will want to modify this entire README to be about YOUR module. Also update the [CONTRIBUTING.md](./CONTRIBUTING.md) file with steps on how to contribute to YOUR module. The [LICENSE](./LICENSE) file is optional if you are NOT publishing to the [Terraform module registry](https://registry.terraform.io/).
## Contributing
See [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

Distributed under the Apache-2.0 License. See [LICENSE](./LICENSE) for more information.

## Contact

Levi - [@shady_cuz](https://twitter.com/shady_cuz)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Release Drafter](https://github.com/release-drafter/release-drafter)
* [pre-commit-terraform]
* [terraform-docs]
* [TFLint]
* [tfsec]
* [checkov]
* [Terratest]
* [Best-README-Template](https://github.com/othneildrew/Best-README-Template)

<br/>
<br/>

### Everything below this line is generated via [terraform-docs]
___
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


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[terraform-shield]: https://img.shields.io/badge/version-Latest%20%7C%200.15%20%7C%200.14%20%7C%200.13%20%7C%200.12-brightgreen?style=for-the-badge&logo=terraform
[tf-version]: ./terraform.tf
[release-url]: https://github.com/DontShaveTheYak/terraform-module-template/releases/latest
[version-shield]: https://img.shields.io/github/v/release/DontShaveTheYak/terraform-module-template?label=latest&style=for-the-badge
[test-shield]: https://img.shields.io/github/workflow/status/DontShaveTheYak/terraform-module-template/Tests?label=Tests&style=for-the-badge
[test-url]: https://github.com/DontShaveTheYak/terraform-module-template/actions?query=workflow%3ATests+branch%3Amain
[codecov-shield]: https://img.shields.io/codecov/c/gh/DontShaveTheYak/terraform-module-template/main?color=green&style=for-the-badge&token=bfF18q99Fl
[codecov-url]: https://codecov.io/gh/DontShaveTheYak/terraform-module-template
[contributors-shield]: https://img.shields.io/github/contributors/DontShaveTheYak/terraform-module-template.svg?style=for-the-badge
[contributors-url]: https://github.com/DontShaveTheYak/terraform-module-template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/DontShaveTheYak/terraform-module-template.svg?style=for-the-badge
[forks-url]: https://github.com/DontShaveTheYak/terraform-module-template/network/members
[stars-shield]: https://img.shields.io/github/stars/DontShaveTheYak/terraform-module-template.svg?style=for-the-badge
[stars-url]: https://github.com/DontShaveTheYak/terraform-module-template/stargazers
[issues-shield]: https://img.shields.io/github/issues/DontShaveTheYak/terraform-module-template.svg?style=for-the-badge
[issues-url]: https://github.com/DontShaveTheYak/terraform-module-template/issues
[license-shield]: https://img.shields.io/github/license/DontShaveTheYak/terraform-module-template.svg?style=for-the-badge
[license-url]: https://github.com/DontShaveTheYak/terraform-module-template/blob/main/LICENSE
[product-screenshot]: images/screenshot.png
[pre-commit-terraform]: https://github.com/antonbabenko/pre-commit-terraform
[terraform-docs]: https://github.com/terraform-docs/terraform-docs/
[TFLint]: https://github.com/terraform-linters/tflint
[tfsec]: https://github.com/aquasecurity/tfsec
[checkov]: https://github.com/bridgecrewio/checkov
[Terratest]: https://github.com/gruntwork-io/terratest