schema "public" {}

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
  column "updated_at" {
    type = timestamp
    null = false
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  column "description" {
    type = varchar(255)
    null = false
  }
  column "country" {
    type = varchar(2)
    null = false
  }
  column "state" {
    type = varchar(100)
    null = false
  }
  column "city" {
    type = varchar(100)
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

/* table "sm_item" {
  schema = schema.public
  column "id" {
    type = integer
    null = false
  }
  column "spot_id" {
    type = integer
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
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
  foreign_key "fk_sm_item_sm_spot" {
    columns     = [column.spot_id]
    ref_columns = [table.sm_spot.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
}

table "sm_order" {
  schema = schema.public
  column "id" {
    type = serial
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
    type = timestamp
    null = false
  }
  column "client_id" {
    type = varchar(10)
    null = false
  }
  column "courier_id" {
    type = varchar(10)
    null = true
  }
  column "status" {
    type = smallint
    null = false
  }
  primary_key {
    columns = [column.id]
  }
}

table "sm_order_item" {
  schema = schema.public
  column "order_id" {
    type = integer
    null = false
  }
  column "item_id" {
    type = integer
    null = false
  }
  primary_key {
    columns = [column.order_id, column.item_id]
  }
  foreign_key "fk_sm_order_item_sm_order" {
    columns     = [column.order_id]
    ref_columns = [table.sm_order.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
  foreign_key "fk_sm_order_item_sm_item" {
    columns     = [column.order_id]
    ref_columns = [table.sm_item.column.id]
    on_delete   = CASCADE
    on_update   = NO_ACTION
  }
} */
