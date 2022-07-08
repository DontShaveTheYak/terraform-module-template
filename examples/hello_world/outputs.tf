output "result" {
  description = "The resul of calling the hello_world module with default args."
  value       = module.hello_world.greeting
}