data "google_project" "default" {}

locals {
  services = [
    "artifactregistry.googleapis.com",
    "run.googleapis.com",
    "secretmanager.googleapis.com"
  ]
}

resource "google_project_service" "default" {
  for_each = toset(local.services)
  service  = each.value
  project  = data.google_project.default.id
}
