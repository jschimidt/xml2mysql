# xml2mysql

Streaming XML profiler and MySQL 8.4 schema generator for very large relational-database XML exports.

## v0.1

The analyzer streams XML rather than loading files into memory. It profiles columns, creates conservative MySQL mappings, and writes profile.json, mapping.json, schema.sql, and an execution report.

Large-file safeguards: 1 MiB buffered sequential reads; no ReadAll/DOM/row accumulation; bounded samples and distinct values; bounded worker pool. Start with 1-2 workers and benchmark storage before increasing concurrency.

## Docker

Copy config.example.yaml to config.yaml, put XML files under data/input, then run:

    docker compose build
    docker compose run --rm xml2mysql

For production, bind-mount the XML directory read-only and output directory read-write. Paths in config.yaml are container paths.

## Scope

v0.1 performs analysis and DDL generation for MySQL 8.4. PKs, FKs and indexes are intentionally not invented. Data import and persistent checkpoint/resume are future work.
