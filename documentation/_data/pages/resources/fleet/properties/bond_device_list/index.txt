---
page_title: "bond_device_list"
subcategory: ""
description: "List of bond devices for this fleet."
xcsh_docs: {"aliases": ["bond device list"], "body_bytes": 1961, "body_sha256": "sha256:2b88fe5ced612d7a39112d50477740ee0c44ad23a01acb707e2cb1101e9311e5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:bond_device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/bond_device_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list:RequiredObjectAttributes:bond_devices", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "sections": [{"aliases": ["bond device list bond devices"], "anchor": "section", "description": "List of bond devices.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:lacp", "type": "conflicts"}, {"anchor": "schema-bond_device_list--bond_devices--devices", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_up_delay", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--name", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}], "schema_path": ["bond_device_list", "bond_devices"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "List of bond devices for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["fleetCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
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

- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/#section)
- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/no_bond_devices/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/): complete subsection reference.

## Next pages

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
