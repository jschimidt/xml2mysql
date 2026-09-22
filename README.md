# xml2mysql

Streaming converter for relational-database XML exports to a MySQL 8.4 dump.

## Purpose

The Carambei job reconstructs the XML export as a temporary database named `carambei_sistema_ids_bkp`. The output is a SQL dump containing **structure and data**. The application does not connect to MySQL; the DBA restores the generated dump manually.

The primary goal is data preservation. Conservative/oversized MySQL types are acceptable when they reduce truncation risk. PKs, FKs and indexes are intentionally not invented.

## Processing model

The converter uses two streaming passes:

1. analyze every XML and infer a conservative schema;
2. stream the XML files again and write batched INSERT statements to the dump.

The XML files and the generated dump are never loaded completely into memory. INSERT batches contain up to 1000 rows.

The main artifact is:

    /data/output/carambei/dump/carambei_sistema_ids_bkp.sql

Auxiliary artifacts remain available under `metadata/`, `schema/` and `reports/`.

The dump includes:

    CREATE DATABASE IF NOT EXISTS `carambei_sistema_ids_bkp`;
    USE `carambei_sistema_ids_bkp`;
    CREATE TABLE ...
    INSERT INTO ... VALUES ...

The execution validates that the number of rows written to the dump equals the number of rows observed during analysis.

## Docker / VM deployment

The VM only needs Git and Docker with Compose support. Go does not need to be installed.

Copy:

    cp config.example.yaml config.yaml

Keep logical paths in `config.yaml`:

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

The default Compose mounts:

    ./data/input  -> /data/input  (read-only)
    ./data/output -> /data/output (read-write)

For a VM, replace the host side of those mounts with the real directories. Do not copy large XML files into the Docker image.

Build and execute:

    docker compose build
    docker compose run --rm xml2mysql version
    docker compose run --rm xml2mysql

## Local reduced test

Put a small set of representative XML files in `data/input/`, run the converter and inspect:

    cat data/output/carambei/reports/analysis.txt
    head -100 data/output/carambei/dump/carambei_sistema_ids_bkp.sql

The report includes `dump_rows`. It must equal `rows`.

## Restore by DBA

The application does not perform this step. The generated file can later be restored by the DBA using their normal MySQL 8.4 process, for example through the MySQL client.

## Scope

Current scope: XML -> complete MySQL 8.4 SQL dump (DDL + data). No MySQL credentials or database connectivity are required by xml2mysql.
