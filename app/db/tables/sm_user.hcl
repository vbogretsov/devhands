table "sm_user" {
  schema = schema.public
  column "id" {
    type = serial
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "email" {
    type = varchar(255)
    null = false
    unique = true
  }
  column "first_name" {
    type = varchar(255)
    null = false
  }
  column "last_name" {
    type = varchar(255)
    null = false
  }
  primary_key {
    columns = [column.id]
  }
}
