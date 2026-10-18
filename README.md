# Gator

Gator is a command-line RSS feed aggregator written in Go. It stores users, feeds, follows, and fetched posts in PostgreSQL.

## Requirements

- [Go](https://go.dev/dl/) (the module declares Go 1.27.2)
- PostgreSQL
- [Goose](https://github.com/pressly/goose) to apply the database migrations in `sql/schema`

## Install

Install the CLI with Go:

```sh
go install github.com/ashraddhansh/gator@latest
```

Make sure Go's binary directory (usually `$(go env GOPATH)/bin`) is on your `PATH` so you can run `gator`.

Install Goose, if you do not already have it:

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Configure PostgreSQL and Gator

Create a PostgreSQL database for Gator. For example, using the `createdb` utility:

```sh
createdb gator
```

Set a PostgreSQL connection URL and apply all migrations from the repository root:

```sh
export DB_URL="postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
goose -dir sql/schema postgres "$DB_URL" up
```

Change the username, password, host, port, and database in `DB_URL` to match your PostgreSQL setup. Goose applies the schema for users, feeds, feed follows, and posts.

Create `~/.gatorconfig.json` (in your home directory) with the same database URL:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

The `db_url` field is required. Gator updates `current_user_name` when you register or log in. Keep the URL in the config in sync with the URL used for migrations.

## Run

After installing the CLI and configuring the database, run commands from any directory:

```sh
gator register alice
```

Registration creates the user and makes that user current. You can also switch to an existing user with `gator login alice`.

## Commands

```text
gator register <username>         Create a user and make them current
gator login <username>            Set the current user
gator users                       List users
gator addfeed <name> <url>        Add an RSS feed and follow it
gator feeds                       List feeds
gator follow <url>                Follow an existing feed
gator following                   List feeds followed by the current user
gator unfollow <url>              Stop following a feed
gator agg <duration>              Fetch feeds continuously at an interval
gator browse [limit]              Show fetched posts (default limit: 2)
gator reset                       Delete all users and their related data
```

For example, add a feed, then start the aggregator in a terminal:

```sh
gator addfeed golang https://go.dev/blog/feed.atom
gator agg 1m
```

`agg` runs continuously; stop it with Ctrl-C. In another terminal, browse the posts:

```sh
gator browse 10
```

`addfeed`, `follow`, `following`, `unfollow`, and `browse` require a current user. `reset` deletes all users and cascades to their feeds, follows, and posts.
