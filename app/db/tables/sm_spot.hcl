table "sm_spot" {
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
    tupe = varchar(255)
    null = false
  }
  column "lat" {
    type = numeric(9, 6)
    null = false
  }
  column "lng" {
    type = numeric(9, 6)
    null = false
  }
  column "raiting" {
    type = float
  }
  primary_key {
    columns = [column.id]
  }
}
