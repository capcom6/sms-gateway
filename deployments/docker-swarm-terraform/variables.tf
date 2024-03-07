variable "swarm-manager-host" {
  type        = string
  sensitive   = true
  description = "Address of swarm manager"
}

variable "app-name" {
  type        = string
  description = "Name of app"
}

variable "app-version" {
  type        = string
  description = "Version of Docker image of app"
  default     = "1.0"
}

variable "app-host" {
  type        = string
  description = "Hostname of app"
}

variable "app-config-b64" {
  type        = string
  description = "Application config file, base64 encoded"
  sensitive   = true
}

variable "app-env-json-b64" {
  type        = string
  description = "Evironment variables, base64 encoded JSON"
  sensitive   = true
}

variable "memory-limit" {
  type        = number
  description = "Memory limit in bytes"
  default     = 32 * 1024 * 1024
}

