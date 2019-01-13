provider "skhema" {
  namespace = "io.skhema.services"
}

data "skhema_api" "users" {
  namespace = "io.skhema.apis.users.v1"
  name = "users-api"
}

resource "skhema_service" "users" {
  name = "users-api"
  api = [
    "${data.skhema_api.users.urn}"
  ]
}
