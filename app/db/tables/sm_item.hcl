table "sm_item" {
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
  column "description" {
    type = varchar(2047)
    null = false
  }
  column "raiting" {
    type = float
  }
  primary_key {
    columns = [column.id]
  }
}
