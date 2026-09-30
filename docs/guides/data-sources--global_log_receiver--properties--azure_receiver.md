---
page_title: "azure_receiver"
subcategory: ""
description: "azure_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2885, "body_sha256": "sha256:d204cb4e96bbe0a5cbaef4e6c07144d295294f0dd1e0c09c10595bb4428ebfa9", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:connection_string", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:filename_options"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--azure_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/azure_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- azure_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Azure Blob Configuration for Global Log Receiver.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [batch](data-sources--global_log_receiver--properties--azure_receiver--batch.md): complete subsection reference.

- [compression](data-sources--global_log_receiver--properties--azure_receiver--compression.md): complete subsection reference.

- [connection_string](data-sources--global_log_receiver--properties--azure_receiver--connection_string.md): complete subsection reference.

<a id="schema-azure_receiver--container_name"></a>

### container_name property

Type: `"string"`. Computed.

Container Name is the name of the container into which logs should be stored.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

- [filename_options](data-sources--global_log_receiver--properties--azure_receiver--filename_options.md): complete subsection reference.

## Next pages

- [azure_receiver.batch](data-sources--global_log_receiver--properties--azure_receiver--batch.md)
- [azure_receiver.compression](data-sources--global_log_receiver--properties--azure_receiver--compression.md)
- [azure_receiver.connection_string](data-sources--global_log_receiver--properties--azure_receiver--connection_string.md)
- [azure_receiver.filename_options](data-sources--global_log_receiver--properties--azure_receiver--filename_options.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
