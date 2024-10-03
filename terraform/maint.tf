resource local_file name {
  filename             = "~/test.txt"
  content = "This is a test file!"
  file_permission      = 0777
  directory_permission = 0777
}
