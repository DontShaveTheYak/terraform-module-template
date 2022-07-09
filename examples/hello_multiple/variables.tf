variable "friends" {
  type        = list(string)
  default     = ["Terraform", "Terratest", "Terraform Weekly"]
  description = "A list of people/things to greet."
}