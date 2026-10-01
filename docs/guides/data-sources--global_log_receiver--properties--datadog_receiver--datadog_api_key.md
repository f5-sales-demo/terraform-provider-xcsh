---
page_title: "datadog_receiver.datadog_api_key"
subcategory: ""
description: "datadog_receiver.datadog_api_key for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1659, "body_sha256": "sha256:cea7acd85c420649c82ca7a938b57ff0b3a5cb9dbfdfd8e65b694efcea173988", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:datadog_api_key", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:datadog_api_key:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:datadog_api_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:datadog_api_key", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["datadog_receiver", "datadog_api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/datadog_receiver/datadog_api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "datadog_receiver.datadog_api_key for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.datadog_api_key

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [datadog_receiver](data-sources--global_log_receiver--properties--datadog_receiver.md)
- datadog_receiver.datadog_api_key

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

- [blindfold_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [datadog_receiver.datadog_api_key.blindfold_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--blindfold_secret_info.md)
- [datadog_receiver.datadog_api_key.clear_secret_info](data-sources--global_log_receiver--properties--datadog_receiver--datadog_api_key--clear_secret_info.md)
- [datadog_receiver](data-sources--global_log_receiver--properties--datadog_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
