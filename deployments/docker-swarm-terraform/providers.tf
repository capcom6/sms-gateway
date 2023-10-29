terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

provider "docker" {
  host     = var.swarm-manager-host
  ssh_opts = ["-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null"]
  # registry_auth {
  #   address  = "cr.selcloud.ru/soft-c"
  #   username = "token"
  #   password = var.registry-password
  # }
}
