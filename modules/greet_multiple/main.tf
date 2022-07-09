locals {
  len      = length(var.names)
  csv_list = slice(var.names, 0, local.len - 1)
  and_list = slice(var.names, local.len - 1, local.len)
  input    = "${join(", ", local.csv_list)} and ${local.and_list[0]}"
}

module "hello" {
  source = "../../"

  name = local.input
}
