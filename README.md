# Scalab

## Installation

All commands will assume you have an `alias dc=docker-compose`

- `dc up -d`
- install migration tool: `dc exec scalab go get -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate/`
- run migrations: `dc exec scalab /bin/sh -c 'set -a; . ./.env; set +a; migrate -database ${DATABASE_URL} -path db/migrations up'`
- `sudo vim /etc/hosts` and add the line `127.0.0.1 scalab.local`
- navigate to [http://scalab.local](http://scalab.local)
