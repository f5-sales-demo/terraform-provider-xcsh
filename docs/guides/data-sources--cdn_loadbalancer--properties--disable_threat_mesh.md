---
page_title: "disable_threat_mesh"
subcategory: "Load Balancing"
description: "disable_threat_mesh for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1257, "body_sha256": "sha256:1409fb648fbb1645dd19754cd663d80df88a32c61c419994e866a7fd54333d8e", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:disable_threat_mesh", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:disable_threat_mesh", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--disable_threat_mesh.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_threat_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/disable_threat_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_threat_mesh for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_threat_mesh

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- disable_threat_mesh

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](data-sources--cdn_loadbalancer--properties--disable_threat_mesh.md#section)
- [enable_threat_mesh](data-sources--cdn_loadbalancer--properties--enable_threat_mesh.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
