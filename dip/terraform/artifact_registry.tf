resource "google_artifact_registry_repository" "cloudrun" {
  format        = "DOCKER"
  location      = "asia-northeast1"
  repository_id = "cloud-run"
}
