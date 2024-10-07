resource "google_sql_database_instance" "default" {
  name             = "finn-server-tutorial"
  database_version = "POSTGRES_16"
  region           = "asia-northeast1"

  settings {
    tier = "db-f1-micro" # 最弱スペック

    # 公開IPを有効化
    ip_configuration {
      ipv4_enabled = true # パブリックIPを有効化
      authorized_networks {
        name  = "all-public" # 任意の名前
        value = "0.0.0.0/0"  # 全てのIPアドレスからの接続を許可
      }
    }
  }
}

resource "google_sql_user" "root" {
  name     = "postgres"
  password = "postgres"
  instance = google_sql_database_instance.default.name
}

resource "google_sql_database" "default" {
  name     = "ft_prod"
  instance = google_sql_database_instance.default.name
}
