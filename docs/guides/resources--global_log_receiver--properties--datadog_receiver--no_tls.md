---
page_title: "datadog_receiver.no_tls"
subcategory: ""
description: "datadog_receiver.no_tls for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1032, "body_sha256": "sha256:d8c7a98834dede687d074ee949479e0fe251c12f5bb58a2ae9f5910e8a2a2f99", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:no_tls", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:no_tls", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "path": "docs/guides/resources--global_log_receiver--properties--datadog_receiver--no_tls.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["datadog_receiver", "no_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/datadog_receiver/no_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "datadog_receiver.no_tls for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.no_tls

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [datadog_receiver](resources--global_log_receiver--properties--datadog_receiver.md)
- datadog_receiver.no_tls

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
no_tls = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [datadog_receiver](resources--global_log_receiver--properties--datadog_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
