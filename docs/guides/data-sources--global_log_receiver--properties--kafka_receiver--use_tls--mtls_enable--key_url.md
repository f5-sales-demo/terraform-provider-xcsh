---
page_title: "kafka_receiver.use_tls.mtls_enable.key_url"
subcategory: ""
description: "kafka_receiver.use_tls.mtls_enable.key_url for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2010, "body_sha256": "sha256:562165da080005b8011a8fda71a493054cd66c9895c3c0fd398ca65756f82e2c", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable:key_url", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable:key_url:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable:key_url", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "path": "docs/guides/data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kafka_receiver", "use_tls", "mtls_enable", "key_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kafka_receiver.use_tls.mtls_enable.key_url for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.use_tls.mtls_enable.key_url

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable.md)
- kafka_receiver.use_tls.mtls_enable.key_url

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md): complete subsection reference.

## Next pages

- [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--blindfold_secret_info.md)
- [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
