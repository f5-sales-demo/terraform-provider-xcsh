---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_lma_region."
xcsh_docs: {"aliases": [], "body_bytes": 11961, "body_sha256": "sha256:0e0b1882666576c39129b97256f82a3e5237deb50d293cc79cdb82c52f061620", "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "xcsh-docs:data-sources:lma_region:properties:elastic_params", "xcsh-docs:data-sources:lma_region:properties:kafka_params"], "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:reference", "parent_id": "xcsh-docs:data-sources:lma_region:fundamentals", "path": "documentation/data-sources/lma_region/properties/index.md", "provider_name": "lma_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_lma_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- Property reference

## Direct properties

- [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/): complete subsection reference.

<a id="schema-country"></a>

### country property

Type: `"string"`. Computed.

Country associated with this LMA region.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

- [elastic_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/elastic_params/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-is_default"></a>

### is_default property

Type: `"bool"`. Computed.

Is Default. Is this the default region.

- [kafka_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/kafka_params/): complete subsection reference.

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
| `access_logs_s3_params` | [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/#section) |
| `access_logs_s3_params.aws_credentials` | [access_logs_s3_params.aws_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/#section) |
| `access_logs_s3_params.aws_credentials.access_key_id` | [access_logs_s3_params.aws_credentials.access_key_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/#schema-access_logs_s3_params--aws_credentials--access_key_id) |
| `access_logs_s3_params.aws_credentials.region` | [access_logs_s3_params.aws_credentials.region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/#schema-access_logs_s3_params--aws_credentials--region) |
| `access_logs_s3_params.aws_credentials.secret_access_key` | [access_logs_s3_params.aws_credentials.secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/#section) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/blindfold_secret_info/#section) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.decryption_provider` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/blindfold_secret_info/#schema-access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info--decryption_provider) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.location` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/blindfold_secret_info/#schema-access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info--location) |
| `access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.store_provider` | [access_logs_s3_params.aws_credentials.secret_access_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/blindfold_secret_info/#schema-access_logs_s3_params--aws_credentials--secret_access_key--blindfold_secret_info--store_provider) |
| `access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info` | [access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/#section) |
| `access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.provider_ref` | [access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/#schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--provider_ref) |
| `access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.url` | [access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/#schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--url) |
| `access_logs_s3_params.bucket` | [access_logs_s3_params.bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/#schema-access_logs_s3_params--bucket) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-annotations) |
| `clickhouse_params` | [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/#section) |
| `clickhouse_params.host` | [clickhouse_params.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/#schema-clickhouse_params--host) |
| `clickhouse_params.password` | [clickhouse_params.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/#section) |
| `clickhouse_params.password.blindfold_secret_info` | [clickhouse_params.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/#section) |
| `clickhouse_params.password.blindfold_secret_info.decryption_provider` | [clickhouse_params.password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/#schema-clickhouse_params--password--blindfold_secret_info--decryption_provider) |
| `clickhouse_params.password.blindfold_secret_info.location` | [clickhouse_params.password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/#schema-clickhouse_params--password--blindfold_secret_info--location) |
| `clickhouse_params.password.blindfold_secret_info.store_provider` | [clickhouse_params.password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/#schema-clickhouse_params--password--blindfold_secret_info--store_provider) |
| `clickhouse_params.password.clear_secret_info` | [clickhouse_params.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/#section) |
| `clickhouse_params.password.clear_secret_info.provider_ref` | [clickhouse_params.password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/#schema-clickhouse_params--password--clear_secret_info--provider_ref) |
| `clickhouse_params.password.clear_secret_info.url` | [clickhouse_params.password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/#schema-clickhouse_params--password--clear_secret_info--url) |
| `clickhouse_params.port` | [clickhouse_params.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/#schema-clickhouse_params--port) |
| `clickhouse_params.user` | [clickhouse_params.user](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/#schema-clickhouse_params--user) |
| `country` | [country](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-country) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-description) |
| `elastic_params` | [elastic_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/elastic_params/#section) |
| `elastic_params.urls` | [elastic_params.urls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/elastic_params/#schema-elastic_params--urls) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-id) |
| `is_default` | [is_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-is_default) |
| `kafka_params` | [kafka_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/kafka_params/#section) |
| `kafka_params.bootstrap_servers` | [kafka_params.bootstrap_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/kafka_params/#schema-kafka_params--bootstrap_servers) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/#schema-namespace) |

## Next pages

- [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/)
- [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/)
- [elastic_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/elastic_params/)
- [kafka_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/kafka_params/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
