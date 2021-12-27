resource "google_cloud_run_service" "front" {
  name     = "blog-front"
  location = var.region

  template {
    spec {
      containers {
        image = "asia-northeast1-docker.pkg.dev/anan-project/blog/front"
        ports = [
          {
            container_port = 80
          }
        ]
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}