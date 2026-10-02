---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address"
subcategory: ""
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["storage static routes storage routes nexthop nexthop address"], "body_bytes": 3246, "body_sha256": "sha256:e63c92df9f4c1594e286c187846b328984c5ccd9dc1a1024b1a52890c3e99ff1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop", "path": "documentation/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0113300103011330-2021221000112230-2120203033320002-0000031323332030-3330320102133210-3331212021003013-3132112030012121-0110211233022031", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address"], "schema_version": 1, "sections": [{"aliases": ["dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "dual_stack"], "syntax": "block", "type": "object"}, {"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "ipv4"], "syntax": "block", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "ipv6"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
```

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

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv6/): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/dual_stack/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv4/)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv6/)
- [storage_static_routes.storage_routes.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
