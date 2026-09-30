---
page_title: "azure_receiver.batch.timeout_seconds_default"
subcategory: ""
description: "azure_receiver.batch.timeout_seconds_default for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1092, "body_sha256": "sha256:8d28ea0ea4af742b3a9c0d612a389c6c95536a337cac0454fdc26e13ff5cfa9b", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch:timeout_seconds_default", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch:timeout_seconds_default", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:batch", "path": "docs/guides/resources--global_log_receiver--properties--azure_receiver--batch--timeout_seconds_default.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_receiver", "batch", "timeout_seconds_default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/batch/timeout_seconds_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_receiver.batch.timeout_seconds_default for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_receiver.batch.timeout_seconds_default

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
- [azure_receiver.batch](resources--global_log_receiver--properties--azure_receiver--batch.md)
- azure_receiver.batch.timeout_seconds_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
timeout_seconds_default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [azure_receiver.batch](resources--global_log_receiver--properties--azure_receiver--batch.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
