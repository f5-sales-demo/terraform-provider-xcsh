---
page_title: "http_receiver.batch.timeout_seconds_default"
subcategory: ""
description: "http_receiver.batch.timeout_seconds_default for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1183, "body_sha256": "sha256:af63587b8031b4f3823cc40d74eb96756997f202a023095e3aa57e4a830cbf14", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch:timeout_seconds_default", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch:timeout_seconds_default", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:batch", "path": "docs/guides/resources--global_log_receiver--properties--http_receiver--batch--timeout_seconds_default.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "batch", "timeout_seconds_default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/batch/timeout_seconds_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.batch.timeout_seconds_default for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.batch.timeout_seconds_default

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [http_receiver](resources--global_log_receiver--properties--http_receiver.md)
- [http_receiver.batch](resources--global_log_receiver--properties--http_receiver--batch.md)
- http_receiver.batch.timeout_seconds_default

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

- [http_receiver.batch](resources--global_log_receiver--properties--http_receiver--batch.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
