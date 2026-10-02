---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address"
subcategory: ""
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["storage static routes storage routes nexthop nexthop address"], "body_bytes": 2867, "body_sha256": "sha256:34fe6723308041a9fb393b49451d7f1c3ff81cec3b2ba0d24d63a7f95620548d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop", "path": "documentation/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312", "registry_path": "docs/guides/data-sources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address"], "schema_version": 1, "sections": [{"aliases": ["dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
