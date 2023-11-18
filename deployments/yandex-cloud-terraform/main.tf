locals {
  sa-roles = toset([
    "container-registry.images.puller",
    "lockbox.payloadViewer",
  ])
}

resource "yandex_iam_service_account" "sa" {
  folder_id = var.folder-id
  name      = "${var.app-name}-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "catgpt-roles" {
  for_each  = local.sa-roles
  folder_id = var.folder-id
  member    = "serviceAccount:${yandex_iam_service_account.sa.id}"
  role      = each.key
}

resource "yandex_serverless_container" "container" {
  folder_id          = var.folder-id
  name               = var.app-name
  memory             = 128
  execution_timeout  = "5s"
  cores              = 1
  core_fraction      = 5
  service_account_id = yandex_iam_service_account.sa.id
  secrets {
    id                   = yandex_lockbox_secret.secret.id
    version_id           = yandex_lockbox_secret_version.secret-version.id
    key                  = "database-host"
    environment_variable = "DATABASE__HOST"
  }
  secrets {
    id                   = yandex_lockbox_secret.secret.id
    version_id           = yandex_lockbox_secret_version.secret-version.id
    key                  = "database-user"
    environment_variable = "DATABASE__USER"
  }
  secrets {
    id                   = yandex_lockbox_secret.secret.id
    version_id           = yandex_lockbox_secret_version.secret-version.id
    key                  = "database-password"
    environment_variable = "DATABASE__PASSWORD"
  }
  secrets {
    id                   = yandex_lockbox_secret.secret.id
    version_id           = yandex_lockbox_secret_version.secret-version.id
    key                  = "database-database"
    environment_variable = "DATABASE__DATABASE"
  }
  secrets {
    id                   = yandex_lockbox_secret.secret.id
    version_id           = yandex_lockbox_secret_version.secret-version.id
    key                  = "goose-dsn"
    environment_variable = "GOOSE_DBSTRING"
  }

  image {
    url = var.image-name
    environment = {
      "HTTP__LISTEN"       = "0.0.0.0:3000"
      "DATABASE__TIMEZONE" = "Europe/Moscow"
    }
  }
}
