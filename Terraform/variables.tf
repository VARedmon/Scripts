variable "filename" {
  default = "~/test.txt"
  type = string
  
}
variable "file-content" {
  type = map
  default = {
    "statement1" = "We love pets!"
    "statement2" = "We love animals!"
  }
}
variable "prefix" {
  default = ["Mr", "Mrs", "Sir"]
  type = list

}
variable "separator" {
  default = "."
  type = string

}
variable "length" {
  default = "1"
  type = string
  
}
