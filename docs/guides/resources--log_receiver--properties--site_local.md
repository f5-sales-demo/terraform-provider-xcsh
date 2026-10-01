---
page_title: "site_local"
subcategory: "Monitoring"
description: "site_local for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 878, "body_sha256": "sha256:08c898ee4d85949fb8bff34d4a337d94e3367d695ae6c52a128d1838d3f38ea9", "canonical_id": "xcsh-docs:resources:log_receiver:properties:site_local", "child_ids": [], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:site_local", "parent_id": "xcsh-docs:resources:log_receiver:reference", "path": "docs/guides/resources--log_receiver--properties--site_local.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_local"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/site_local/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_local for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_local

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Property reference](resources--log_receiver--reference.md)
- site_local

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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
site_local {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--log_receiver--reference.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
