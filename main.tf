

module "say_hello" {
  source = "./modules/say_hello"

  name = var.name

}

module "make_exciting" {
  source = "./modules/make_exciting"

  text = module.say_hello.result
}