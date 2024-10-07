provider "google" {
  project = "nobuz-437711"
}

terraform {
  required_providers {
    google = {
      version = "~> 5.3.0"
    }
  }
}
