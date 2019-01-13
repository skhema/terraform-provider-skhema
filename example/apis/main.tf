provider "skhema" {
  namespace = "io.skhema.apis.users.v1"
}

data "skhema_type" "user" {
  namespace = "io.skhema.types.users.v1"
  name = "User"
}

resource "skhema_api" "users" {
  name = "users-api"
  /*consumes = [
    "application/json"
  ]

  produces = [
    "application/json"
  ]*/

  operation {
    name = "create_user"
    path = "/users"
    method = "post"

    consume {
      format = "json"

      type = "record"
      schema = "${data.skhema_type.user.schema}"
    }

    produce {
      status = "201"

      format = "json"
      type = "record"
      schema = "${data.skhema_type.user.schema}"
    }
  }

  operation {
    name = "list_users"
    path = "/users"
    method = "get"

    param {
      name = "limit"
      segment = "query"
      type = "string"
    }

    produce {
      status = "200"

      format = "json"
      type = "array"
      schema = "${data.skhema_type.user.schema}"
    }
  }
}
