provider "skhema" {
  namespace = "io.skhema.types.users.v1"
}

resource "skhema_type" "user" {
  // doc = ""
  name = "User"
  type = "record"

  field {
    // doc = ""
    name = "id"
    type = "string"
    // required = true
    // optional = true
    // computed = true - read-only field
  }

  field {
    name = "email"
    type = "string"
  }

  field {
    name = "name"
    type = "string"
  }

  field {
    name = "deleted"
    type = "string"
  }

  field {
    name = "updated"
    type = "string"
  }
}

output "urn" {
  value = "${skhema_type.user.urn}"
}
