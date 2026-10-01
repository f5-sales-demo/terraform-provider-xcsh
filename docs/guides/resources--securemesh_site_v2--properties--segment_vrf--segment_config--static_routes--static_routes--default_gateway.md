---
page_title: "segment_vrf.segment_config.static_routes.static_routes.default_gateway"
subcategory: ""
description: "segment_vrf.segment_config.static_routes.static_routes.default_gateway for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1628, "body_sha256": "sha256:e307e043c400620f60014391d421018c29a47041f089ac1ec8a8b9ddc9b26c5b", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:default_gateway", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:default_gateway", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "path": "docs/guides/resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_routes--static_routes--default_gateway.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes", "default_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/default_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_vrf.segment_config.static_routes.static_routes.default_gateway for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_routes.static_routes.default_gateway

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [segment_vrf](resources--securemesh_site_v2--properties--segment_vrf.md)
- [segment_vrf.segment_config](resources--securemesh_site_v2--properties--segment_vrf--segment_config.md)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_routes.md)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_routes--static_routes.md)
- segment_vrf.segment_config.static_routes.static_routes.default_gateway

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_gateway = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_routes--static_routes.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
