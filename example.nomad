job "danimunf-pengajuan-pinjaman" {
  datacenters = ["dc1"]

  group "cache" {
    network {
      port "db" {
        to = 1323
      }
    }

    task "redis" {
      driver = "docker"

      config {
        image = "danimunf/pengajuan-pinjaman:latest"

        ports = ["db"]
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
