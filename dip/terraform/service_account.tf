locals {
  cloud_run_roles = [
    "roles/owner",
  ]
}

resource "google_service_account" "cloud_run" {
  account_id   = "cloud-run-sa"
  display_name = "Cloud Run(nobuz) Service Account"
}

resource "google_project_iam_member" "cloud_run" {
  for_each = toset(local.cloud_run_roles)
  role     = each.value
  member   = "serviceAccount:${google_service_account.cloud_run.email}"
  project  = data.google_project.default.id
}
