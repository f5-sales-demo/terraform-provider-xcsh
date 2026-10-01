---
page_title: "qradar_receiver.use_tls.mtls_disabled"
subcategory: ""
description: "qradar_receiver.use_tls.mtls_disabled for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1181, "body_sha256": "sha256:bdbb633b5625070f610722c7f27f20a0a3c394252588646b7d85b62192896f98", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_disabled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls", "path": "docs/guides/resources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_disabled.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["qradar_receiver", "use_tls", "mtls_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "qradar_receiver.use_tls.mtls_disabled for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.use_tls.mtls_disabled

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [qradar_receiver](resources--global_log_receiver--properties--qradar_receiver.md)
- [qradar_receiver.use_tls](resources--global_log_receiver--properties--qradar_receiver--use_tls.md)
- qradar_receiver.use_tls.mtls_disabled

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
mtls_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [qradar_receiver.use_tls](resources--global_log_receiver--properties--qradar_receiver--use_tls.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
