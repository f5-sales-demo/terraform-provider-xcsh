---
page_title: "storage_static_routes.storage_routes.nexthop"
subcategory: ""
description: "storage_static_routes.storage_routes.nexthop for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2912, "body_sha256": "sha256:db14e841a37a41fdafbd45fbd918c85f5b1e86a10cea22958e0e150dce45c6ff", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:interface", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes", "path": "docs/guides/data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.nexthop for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_static_routes.storage_routes.nexthop

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_static_routes](data-sources--fleet--properties--storage_static_routes.md)
- [storage_static_routes.storage_routes](data-sources--fleet--properties--storage_static_routes--storage_routes.md)
- storage_static_routes.storage_routes.nexthop

<a id="section"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

- [interface](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md): complete subsection reference.

- [nexthop_address](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address.md): complete subsection reference.

<a id="schema-storage_static_routes--storage_routes--nexthop--type"></a>

### type property

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [storage_static_routes.storage_routes.nexthop.interface](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address.md)
- [storage_static_routes.storage_routes](data-sources--fleet--properties--storage_static_routes--storage_routes.md)
- [xcsh_fleet](../data-sources/fleet.md)
