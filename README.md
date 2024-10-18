# launchpad-go-rest

A template for back end REST API server, written in Go.

# Setting Up

Run all of the needed systems in your local by using the following command:

```
docker compose up -d
```

or if you have docker-compose installed (the old version of `docker compose`):

```
docker-compose up -d
```

# Server Start

After you have all the systems booted up, you can start the server by using the following command:

```
make start-server
```

You can use the following command to start the cronjob instance:

```
make start-cron
```

You can use the following command to start the queue listener instance:

```
make start-queue
```

# Development Tools

There are some out-of-the-box tools that you can use to streamline your development process:

1. You can access port `:1323` with `/swagger` in your browser to utilize the swagger documentation (example: `localhost:1323/swagger`).
2. You can access port `:8080` in your browser to access the message queue dashboard (asynqmon) (example: `localhost/8080`).

# Migrations

This template supports migration via golang-migration, you can use the following command to create a new up and down migration files:

```
make migration-new name=your_migration_name_here
```

To run your migration, you can use the following command:

```
make migration-up database=<postgres-dsn-here>
```

To revert all of your migration, you can use the following command:

```
make migration-down database=<postgres-dsn-here>
```
