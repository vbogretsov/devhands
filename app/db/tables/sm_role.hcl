table "sm_role" {
  schema = schema.public
  column "id" {
    type = serial
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  primary_key {
    columns = [column.id]
  }
}
