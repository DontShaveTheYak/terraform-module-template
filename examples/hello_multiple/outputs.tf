output "result" {
  description = "The resul of calling the hello_world submodule greet_multiple."
  value       = module.greet_multiple.greeting
}