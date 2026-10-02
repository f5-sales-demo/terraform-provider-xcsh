---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_lma_region."
xcsh_docs: {"aliases": ["lma region"], "body_bytes": 12060, "body_sha256": "sha256:2fb54283b9397a9e77f25342fbf5310e25ee016e23e11ea8644d45d995454afb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "xcsh-docs:data-sources:lma_region:properties:elastic_params", "xcsh-docs:data-sources:lma_region:properties:kafka_params"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:reference", "parent_id": "xcsh-docs:data-sources:lma_region:fundamentals", "path": "documentation/data-sources/lma_region/properties/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0023210031100010-2112103013031230-3231331013203233-3102020222122202-0111202222003331-1021220322200330-2033303210112210-1133233201200122", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["access logs s3 params"], "anchor": "section", "description": "Configuration parameter for access logs s3 params.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_logs_s3_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["clickhouse params"], "anchor": "section", "description": "Configuration parameter for clickhouse params.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["clickhouse_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["country"], "anchor": "schema-country", "description": "Country associated with this LMA region.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["country"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["elastic params"], "anchor": "section", "description": "Configuration parameter for elastic params.", "document_id": "xcsh-docs:data-sources:lma_region:properties:elastic_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["elastic_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["is default"], "anchor": "schema-is_default", "description": "Is Default. Is this the default region.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["is_default"], "syntax": "attribute", "type": "bool"}, {"aliases": ["kafka params"], "anchor": "section", "description": "Configuration parameter for kafka params.", "document_id": "xcsh-docs:data-sources:lma_region:properties:kafka_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["kafka_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the LmaRegion to look up.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the LmaRegion.", "document_id": "xcsh-docs:data-sources:lma_region:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_lma_region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
