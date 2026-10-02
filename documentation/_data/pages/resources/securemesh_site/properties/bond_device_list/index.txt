---
page_title: "bond_device_list"
subcategory: ""
description: "List of bond devices for this fleet."
xcsh_docs: {"aliases": ["bond device list"], "body_bytes": 2061, "body_sha256": "sha256:a54f083d57b7cee43562e0cd0bfba648d2b29e4ab21cd4e38b22bb778fe4b680", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list", "parent_id": "xcsh-docs:resources:securemesh_site:reference", "path": "documentation/resources/securemesh_site/properties/bond_device_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2310121302213102-1031133011333231-3013111001301100-3012112113020110-3203213233333032-3113030101223330-0221130332031101-3133030202011232", "registry_path": "docs/guides/resources--securemesh_site--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list:RequiredObjectAttributes:bond_devices", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "sections": [{"aliases": ["bond devices"], "anchor": "section", "description": "List of bond devices.", "document_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices:lacp", "type": "conflicts"}, {"anchor": "schema-bond_device_list--bond_devices--devices", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_up_delay", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--name", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "type": "requires"}], "schema_path": ["bond_device_list", "bond_devices"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of bond devices for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
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

- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/#section)
- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/no_bond_devices/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/): complete subsection reference.

## Next pages

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
