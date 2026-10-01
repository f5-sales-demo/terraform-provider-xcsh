---
page_title: "s3_receiver.compression.compression_default"
subcategory: ""
description: "s3_receiver.compression.compression_default for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1221, "body_sha256": "sha256:284f613d76e017ebb6d9b5d39fe3387212f75960ada550773e007d8825750d63", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:compression:compression_default", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:compression:compression_default", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:s3_receiver:compression", "path": "docs/guides/resources--global_log_receiver--properties--s3_receiver--compression--compression_default.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["s3_receiver", "compression", "compression_default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/s3_receiver/compression/compression_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "s3_receiver.compression.compression_default for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver.compression.compression_default

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [s3_receiver](resources--global_log_receiver--properties--s3_receiver.md)
- [s3_receiver.compression](resources--global_log_receiver--properties--s3_receiver--compression.md)
- s3_receiver.compression.compression_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [s3_receiver.compression](resources--global_log_receiver--properties--s3_receiver--compression.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
