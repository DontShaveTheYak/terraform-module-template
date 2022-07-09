module "greet_multiple" {
  source = "../../modules/greet_multiple"

  names = var.friends
}