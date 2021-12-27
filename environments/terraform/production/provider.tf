terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  #credentials = file("<NAME>.json")

  project = var.project
  region  = var.region
  zone    = var.zone
}

