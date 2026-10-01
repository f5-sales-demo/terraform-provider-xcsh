---
page_title: "static_routes.default_gateway"
subcategory: "Networking"
description: "static_routes.default_gateway for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 996, "body_sha256": "sha256:3789ae0dd7e31f7107585ef4081ff53aa0986efad63f6a916a8f142be2f1bffb", "canonical_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:default_gateway", "child_ids": [], "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:default_gateway", "parent_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes", "path": "docs/guides/data-sources--virtual_network--properties--static_routes--default_gateway.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["static_routes", "default_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/properties/static_routes/default_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "static_routes.default_gateway for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# static_routes.default_gateway

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md)
- [Property reference](data-sources--virtual_network--reference.md)
- [static_routes](data-sources--virtual_network--properties--static_routes.md)
- static_routes.default_gateway

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [static_routes](data-sources--virtual_network--properties--static_routes.md)
- [xcsh_virtual_network](../data-sources/virtual_network.md)
