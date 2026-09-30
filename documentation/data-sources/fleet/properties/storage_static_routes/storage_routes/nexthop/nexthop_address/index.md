---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address"
subcategory: ""
description: "storage_static_routes.storage_routes.nexthop.nexthop_address for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2768, "body_sha256": "sha256:e0be0519a0f332ae3342f9c9b777be5732218882431fc15989ba24790dc3fb3c", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "path": "documentation/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.nexthop.nexthop_address for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_static_routes.storage_routes.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
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

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv6/): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv4/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv6/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
