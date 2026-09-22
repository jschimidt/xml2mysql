# xml2mysql

Streaming XML profiler and MySQL 8.4 schema generator for very large relational-database XML exports.

## Purpose

The current job reconstructs the Carambei XML export into a temporary MySQL database named `carambei_sistema_ids_bkp`. The database is intended as a migration source for the ESUS development team, not as a permanent production schema.

The primary goal is data preservation. The generated schema is deliberately conservative: oversized types are acceptable when they reduce the risk of truncation or data loss. PKs, FKs and indexes are not invented.

## v0.1

The analyzer streams XML rather than loading files into memory. It profiles columns, creates conservative MySQL mappings, and writes `profile.json`, `mapping.json`, `schema.sql`, and an execution report.

Large-file safeguards: 1 MiB buffered sequential reads; no ReadAll/DOM/row accumulation; bounded samples and distinct values; bounded worker pool. Start with 1-2 workers and benchmark storage before increasing concurrency.

## Docker / VM deployment

The VM only needs Git and Docker with Compose support. Go does not need to be installed on the VM.

Copy `config.example.yaml` to `config.yaml`. Configure the host directories in `compose.yaml` or replace the example bind mounts with the VM paths:

    <host XML directory>:/data/input:ro
    <host output directory>:/data/output
    ./config.yaml:/app/config.yaml:ro

Keep application paths in `config.yaml` as container paths:

    input:
      directory: /data/input

    output:
      directory: /data/output

    target:
      database: mysql
      version: "8.4"
      database_name: carambei_sistema_ids_bkp
      engine: InnoDB
      charset: utf8mb4

Build and run:

    docker compose build
    docker compose run --rm xml2mysql version
    docker compose run --rm xml2mysql

The generated DDL is written to:

    data/output/carambei/schema/schema.sql

It includes `CREATE DATABASE IF NOT EXISTS carambei_sistema_ids_bkp` and `USE carambei_sistema_ids_bkp`.

For production/VM execution, bind-mount the XML directory read-only and the output directory read-write. Do not copy the XML files into the Docker image.

## Scope

v0.1 performs streaming analysis and DDL generation for MySQL 8.4. It does **not yet import XML rows into MySQL**. The generated `schema.sql` is the handoff for the DBA/infrastructure step. Streaming data import and persistent checkpoint/resume are the next phase.
