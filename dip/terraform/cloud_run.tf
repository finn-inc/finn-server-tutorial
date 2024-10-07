resource "google_cloud_run_service" "default" {
  name     = "api"
  location = "asia-northeast1"
  template {
    spec {
      containers {
        name  = "api"
        image = "gcr.io/cloudrun/hello"
        volume_mounts {
          mount_path = "/usr/src/app"
          name       = "secret-env"
        }
        ports {
          container_port = 8080
        }
	env {
		name="DATABASE_URL"
		value="postgresql://postgres:postgres@34.84.228.150/ft_prod?sslmode=disable"
		}
      }

      service_account_name = google_service_account.cloud_run.email
      volumes {
        name = "secret-env"
        secret {
          secret_name = google_secret_manager_secret.api_secret_env.secret_id
          items {
            key  = "latest"
            path = "./.env"
          }
        }
      }
    }
  }

  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
      template[0].spec[0].containers[0].image,
    ]
  }

  depends_on = [google_project_service.default]
}

resource "google_cloud_run_service_iam_binding" "default" {
  location = google_cloud_run_service.default.location
  project  = google_cloud_run_service.default.project
  service  = google_cloud_run_service.default.name
  role     = "roles/run.invoker"
  members = [
    "allUsers"
  ]
  depends_on = [google_cloud_run_service.default]
}
