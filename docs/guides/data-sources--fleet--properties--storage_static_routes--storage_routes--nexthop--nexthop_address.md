---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address"
subcategory: ""
description: "storage_static_routes.storage_routes.nexthop.nexthop_address for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2231, "body_sha256": "sha256:2fcab51bf0ac7909817e92e8b7dae7b6c2fdfdd57d4a9b1c22a53909192b5c94", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "path": "docs/guides/data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.nexthop.nexthop_address for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_static_routes](data-sources--fleet--properties--storage_static_routes.md)
- [storage_static_routes.storage_routes](data-sources--fleet--properties--storage_static_routes--storage_routes.md)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop.md)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="section"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

## Direct properties

- [dual_stack](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack.md): complete subsection reference.

- [ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4.md): complete subsection reference.

- [ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv6.md): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack.md)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4.md)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv6.md)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop.md)
- [xcsh_fleet](../data-sources/fleet.md)
