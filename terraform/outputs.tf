output "filename" {
  description = "The path of the generated file."
  value       = local_file.test.filename
}

output "pet_name" {
  description = "The generated random pet name."
  value       = random_pet.my_pet.id
}

output "file_content_statement1" {
  description = "The first statement written to the file."
  value       = local_file.test.content
}
