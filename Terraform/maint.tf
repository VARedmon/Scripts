resource local_file name {
  sensitive_content = ""
  filename             = "${path.module}/files/outputfile"
  file_permission      = 0777
  directory_permission = 0777
}
