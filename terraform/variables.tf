variable "filename" {
  description = "The output file path."
  type        = string
  default     = "~/test.txt"
}

variable "file_content" {
  description = "Map of file content values."
  type        = map(string)
  default = {
    statement1 = "We love pets!"
    statement2 = "We love animals!"
  }
}

variable "prefix" {
  description = "Prefix choices for the random pet name."
  type        = list(string)
  default     = ["Mr", "Mrs", "Sir"]
}

variable "separator" {
  description = "Separator used in the random pet name."
  type        = string
  default     = "."
}

variable "length" {
  description = "Length of the random pet name suffix."
  type        = number
  default     = 1
}
