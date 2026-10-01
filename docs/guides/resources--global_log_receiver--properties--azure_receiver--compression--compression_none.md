---
page_title: "azure_receiver.compression.compression_none"
subcategory: ""
description: "azure_receiver.compression.compression_none for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1233, "body_sha256": "sha256:b5db0636dbe7184bb988222a1db198664f9d280b3a9584fec82f967a612f8e07", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression", "path": "docs/guides/resources--global_log_receiver--properties--azure_receiver--compression--compression_none.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_receiver", "compression", "compression_none"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/compression/compression_none/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_receiver.compression.compression_none for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.compression.compression_none

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
- [azure_receiver.compression](resources--global_log_receiver--properties--azure_receiver--compression.md)
- azure_receiver.compression.compression_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
compression_none = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [azure_receiver.compression](resources--global_log_receiver--properties--azure_receiver--compression.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
