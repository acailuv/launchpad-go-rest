# launchpad-go-rest

A template for back end REST API server, written in Go.

# Setting Up

Run postgres in your local by using the following command:

```
docker compose up -d postgres
```

or if you have docker-compose installed (the old version of `docker compose`):

```
docker-compose up -d postgres
```

# Server Start

After you have the database booted up, you can start the server by using the following command:

```
make start-server
```

You can use the following command to start the cronjob instance:

```
make start-cron
```

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

# Template Roadmap

- [x] Unit Tests
- [ ] Redis Support
- [ ] Message Broker Support (RMQ/Kafka)
