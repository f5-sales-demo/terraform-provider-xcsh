---
page_title: "bond_device_list"
subcategory: ""
description: "List of bond devices for this fleet."
xcsh_docs: {"aliases": ["bond device list"], "body_bytes": 2051, "body_sha256": "sha256:6e7b2ce8b942e319d0392676cde1506df4a8cd99df07766ed081b9638b6361ed", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/bond_device_list/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3030103000231013-1102103111323200-0331203022323213-3221021030201010-2113303311303120-3003303202322133-2100131233333213-2032320213020130", "registry_path": "docs/guides/resources--voltstack_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list:RequiredObjectAttributes:bond_devices", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "sections": [{"aliases": ["bond device list bond devices"], "anchor": "section", "description": "List of bond devices.", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:lacp", "type": "conflicts"}, {"anchor": "schema-bond_device_list--bond_devices--devices", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_up_delay", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--name", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "type": "requires"}], "schema_path": ["bond_device_list", "bond_devices"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of bond devices for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- bond_device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bond_devices")}
```

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

- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/#section)
- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/no_bond_devices/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/bond_devices/): complete subsection reference.

## Next pages

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/bond_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
