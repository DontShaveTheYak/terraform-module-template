output "result" {
  description = "The resul of calling the hello_world module by passing the name arg."
  value       = module.hello_world.greeting
}