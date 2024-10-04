resource local_file "test" {
  filename = var.filename
  content = var.file-content["statement1"]
  
}
resource "random_pet" "my_pet" {
  prefix = var.prefix[0]
  separator = var.separator
  length = var.length
}