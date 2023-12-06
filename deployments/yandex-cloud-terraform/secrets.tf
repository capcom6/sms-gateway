resource "yandex_lockbox_secret" "secret" {
  name        = "${var.app-name}-secrets"
  description = ""
  folder_id   = var.folder-id
  labels = {
    app = var.app-name
  }
}

resource "yandex_lockbox_secret_version" "secret-version" {
  secret_id = yandex_lockbox_secret.secret.id
  entries {
    key        = "database-host"
    text_value = var.env["database-host"]
  }
  entries {
    key        = "database-user"
    text_value = var.env["database-user"]
  }
  entries {
    key        = "database-password"
    text_value = var.env["database-password"]
  }
  entries {
    key        = "database-database"
    text_value = var.env["database-database"]
  }
}
