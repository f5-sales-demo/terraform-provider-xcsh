---
page_title: "bond_device_list.bond_devices.active_backup"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["bond device list bond devices active backup"], "body_bytes": 1168, "body_sha256": "sha256:d3179b37458919be6a5480b855b7d3fc1daae5fcb27e18f4ffdbb0c060410a06", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:active_backup", "parent_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "path": "documentation/resources/fleet/properties/bond_device_list/bond_devices/active_backup/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2320220031020231-2330123331333300-0121030331313323-1311323222233031-3201020003201011-3111210101112123-2301231002321223-1311303002311002", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list", "bond_devices", "active_backup"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/bond_device_list/bond_devices/active_backup/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["fleetCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list.bond_devices.active_backup

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/)
- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/)
- bond_device_list.bond_devices.active_backup

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

Additional upstream details:

This can be used for messages where no values are needed.

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
active_backup = {}
```

This is an empty object or choice marker. It has no direct properties.
