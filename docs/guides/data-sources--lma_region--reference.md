---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_lma_region."
xcsh_docs: {"aliases": [], "body_bytes": 9688, "body_sha256": "sha256:47ad23be9d7dd9568fb54f07bdbce9262dbfb04a5108648b2c95be6fded0b4b0", "canonical_id": "xcsh-docs:data-sources:lma_region:reference", "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "xcsh-docs:data-sources:lma_region:properties:elastic_params", "xcsh-docs:data-sources:lma_region:properties:kafka_params"], "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:reference", "parent_id": "xcsh-docs:data-sources:lma_region:fundamentals", "path": "docs/guides/data-sources--lma_region--reference.md", "provider_name": "lma_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_lma_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_lma_region](../data-sources/lma_region.md)
- Property reference

## Direct properties

- [access_logs_s3_params](data-sources--lma_region--properties--access_logs_s3_params.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [clickhouse_params](data-sources--lma_region--properties--clickhouse_params.md): complete subsection reference.

<a id="schema-country"></a>

### country property

Type: `"string"`. Computed.

Country associated with this LMA region.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

- [elastic_params](data-sources--lma_region--properties--elastic_params.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-is_default"></a>

### is_default property

Type: `"bool"`. Computed.

Is Default. Is this the default region.

- [kafka_params](data-sources--lma_region--properties--kafka_params.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the LmaRegion to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the LmaRegion.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_logs_s3_params` | [access_logs_s3_params](data-sources--lma_region--properties--access_logs_s3_params.md#section) |
| `access_logs_s3_params.aws_credentials` | [access_logs_s3_params.aws_credentials](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials.md#section) |
| `access_logs_s3_params.aws_credentials.access_key_id` | [access_logs_s3_params.aws_credentials.access_key_id](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials.md#schema-access_logs_s3_params--aws_credentials--access_key_id) |
| `access_logs_s3_params.aws_credentials.region` | [access_logs_s3_params.aws_credentials.region](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials.md#schema-access_logs_s3_params--aws_credentials--region) |
| `access_logs_s3_params.aws_credentials.secret_access_key` | [access_logs_s3_params.aws_credentials.secret_access_key](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key.md#section) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info.md#section) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.decryption_provider` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.decryption_provider](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info.md#schema-access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info--decryption_provider) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.location` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.location](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info.md#schema-access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info--location) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.store_provider` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.store_provider](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info.md#schema-access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info--store_provider) |
| `access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info` | [access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info.md#section) |
| `access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.provider_ref` | [access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.provider_ref](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info.md#schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--provider_ref) |
| `access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.url` | [access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.url](data-sources--lma_region--properties--access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info.md#schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--url) |
| `access_logs_s3_params.bucket` | [access_logs_s3_params.bucket](data-sources--lma_region--properties--access_logs_s3_params.md#schema-access_logs_s3_params--bucket) |
| `annotations` | [annotations](data-sources--lma_region--reference.md#schema-annotations) |
| `clickhouse_params` | [clickhouse_params](data-sources--lma_region--properties--clickhouse_params.md#section) |
| `clickhouse_params.host` | [clickhouse_params.host](data-sources--lma_region--properties--clickhouse_params.md#schema-clickhouse_params--host) |
| `clickhouse_params.password` | [clickhouse_params.password](data-sources--lma_region--properties--clickhouse_params--password.md#section) |
| `clickhouse_params.password.blindfold_secret_info` | [clickhouse_params.password.blindfold_secret_info](data-sources--lma_region--properties--clickhouse_params--password--blindfold_secret_info.md#section) |
| `clickhouse_params.password.blindfold_secret_info.decryption_provider` | [clickhouse_params.password.blindfold_secret_info.decryption_provider](data-sources--lma_region--properties--clickhouse_params--password--blindfold_secret_info.md#schema-clickhouse_params--password--blindfold_secret_info--decryption_provider) |
| `clickhouse_params.password.blindfold_secret_info.location` | [clickhouse_params.password.blindfold_secret_info.location](data-sources--lma_region--properties--clickhouse_params--password--blindfold_secret_info.md#schema-clickhouse_params--password--blindfold_secret_info--location) |
| `clickhouse_params.password.blindfold_secret_info.store_provider` | [clickhouse_params.password.blindfold_secret_info.store_provider](data-sources--lma_region--properties--clickhouse_params--password--blindfold_secret_info.md#schema-clickhouse_params--password--blindfold_secret_info--store_provider) |
| `clickhouse_params.password.clear_secret_info` | [clickhouse_params.password.clear_secret_info](data-sources--lma_region--properties--clickhouse_params--password--clear_secret_info.md#section) |
| `clickhouse_params.password.clear_secret_info.provider_ref` | [clickhouse_params.password.clear_secret_info.provider_ref](data-sources--lma_region--properties--clickhouse_params--password--clear_secret_info.md#schema-clickhouse_params--password--clear_secret_info--provider_ref) |
| `clickhouse_params.password.clear_secret_info.url` | [clickhouse_params.password.clear_secret_info.url](data-sources--lma_region--properties--clickhouse_params--password--clear_secret_info.md#schema-clickhouse_params--password--clear_secret_info--url) |
| `clickhouse_params.port` | [clickhouse_params.port](data-sources--lma_region--properties--clickhouse_params.md#schema-clickhouse_params--port) |
| `clickhouse_params.user` | [clickhouse_params.user](data-sources--lma_region--properties--clickhouse_params.md#schema-clickhouse_params--user) |
| `country` | [country](data-sources--lma_region--reference.md#schema-country) |
| `description` | [description](data-sources--lma_region--reference.md#schema-description) |
| `elastic_params` | [elastic_params](data-sources--lma_region--properties--elastic_params.md#section) |
| `elastic_params.urls` | [elastic_params.urls](data-sources--lma_region--properties--elastic_params.md#schema-elastic_params--urls) |
| `id` | [id](data-sources--lma_region--reference.md#schema-id) |
| `is_default` | [is_default](data-sources--lma_region--reference.md#schema-is_default) |
| `kafka_params` | [kafka_params](data-sources--lma_region--properties--kafka_params.md#section) |
| `kafka_params.bootstrap_servers` | [kafka_params.bootstrap_servers](data-sources--lma_region--properties--kafka_params.md#schema-kafka_params--bootstrap_servers) |
| `labels` | [labels](data-sources--lma_region--reference.md#schema-labels) |
| `name` | [name](data-sources--lma_region--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--lma_region--reference.md#schema-namespace) |

## Next pages

- [access_logs_s3_params](data-sources--lma_region--properties--access_logs_s3_params.md)
- [clickhouse_params](data-sources--lma_region--properties--clickhouse_params.md)
- [elastic_params](data-sources--lma_region--properties--elastic_params.md)
- [kafka_params](data-sources--lma_region--properties--kafka_params.md)
- [xcsh_lma_region](../data-sources/lma_region.md)
