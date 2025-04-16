set dotenv-path := ".env"

export PGHOST := "127.0.0.1"
export PGPORT := "40000"
export PGUSER := `echo $POSTGRES_USER`
export PGPASSWORD := `echo $POSTGRES_PASSWORD`

psql *FLAGS:
    @docker compose exec -e PGUSER=$PGUSER -e PGPASSWORD=$PGPASSWORD postgres psql {{FLAGS}}

[working-directory: 'apps/app']
app-dbdiff:
    @atlas migrate diff --env main --var dbname=delivery

[working-directory: 'apps/app']
app-dbmigrate:
    @atlas migrate apply --env main --var dbname=delivery
