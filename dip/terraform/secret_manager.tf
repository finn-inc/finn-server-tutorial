resource "google_secret_manager_secret" "api_secret_env" {
  secret_id = "api_secret_env"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "api_secret_env" {
  secret = google_secret_manager_secret.api_secret_env.id

  secret_data = "first add"

  lifecycle {
    ignore_changes = [
      secret_data
    ]
  }
}
