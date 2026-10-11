---
page_title: "bond_device_list"
subcategory: ""
description: "List of bond devices for this fleet."
xcsh_docs: {"aliases": ["bond device list"], "body_bytes": 1568, "body_sha256": "sha256:54bc6a07a11e9e473f32888dc7d03ca1601a888c4d653426a15d7aa3a190cf0a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:bond_device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/bond_device_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331", "registry_path": "docs/guides/resources--fleet--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list:RequiredObjectAttributes:bond_devices", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "sections": [{"aliases": ["bond device list bond devices"], "anchor": "section", "description": "List of bond devices.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:lacp", "type": "conflicts"}, {"anchor": "schema-bond_device_list--bond_devices--devices", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_up_delay", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--name", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}], "schema_path": ["bond_device_list", "bond_devices"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of bond devices for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fleetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
