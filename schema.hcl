table "posts" {
  schema = schema.public

  column "id" {
    null = false
    type = bigint
  }
  column "title" {
    null = false
    type = character_varying
  }
  column "body" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
}

schema "public" {
  comment = "standard public schema"
}
