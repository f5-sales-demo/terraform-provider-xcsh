---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack"
subcategory: ""
description: "DualStackAddressType represents both IPv4 and IPv6 together."
xcsh_docs: {"aliases": ["storage static routes storage routes nexthop nexthop address dual stack"], "body_bytes": 2784, "body_sha256": "sha256:f85bb5304d5516bc58f15bca17816456b1e8f51baa3aa436f35bce454b86b61f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack:ipv4", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack:ipv6"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "path": "documentation/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3102121312220313-0113030102112222-1312112312222122-2222211021110113-3333023030230230-0211333200121331-3222310113021231-1001121303013130", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack"], "schema_version": 1, "sections": [{"aliases": ["storage static routes storage routes nexthop nexthop address dual stack ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack:ipv4", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack", "ipv4"], "syntax": "block", "type": "object"}, {"aliases": ["storage static routes storage routes nexthop nexthop address dual stack ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack:ipv6", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "DualStackAddressType represents both IPv4 and IPv6 together.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/ipv6/): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/ipv4/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/ipv6/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
