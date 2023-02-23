terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

provider "google" {
  credentials = file("./credentials.json")

  project = var.project_id
  region  = var.region
  zone    = "us-central1-c"
}

provider "google-beta" {
  project = var.project_id
  region  = var.region
}