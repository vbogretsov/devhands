variable "dbname" {
  type = string
  default = "delivery"
}

variable "user" {
  type = string
  default = getenv("PGUSER")
}

variable "password" {
  type = string
  default = getenv("PGPASSWORD")
}

variable "host" {
  type = string
  default = "localhost"
}

variable "port" {
  type = number
  default = 40000
}

env "main" {
  dir = "file://db/migrations"
  dev = "docker://postgres/17/dev?search_path=public"
  url = format("postgres://%s:%s@%s:%d/%s?sslmode=disable", var.user, var.password, var.host, var.port, var.dbname)
  src = "file://db/schema.hcl"
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
