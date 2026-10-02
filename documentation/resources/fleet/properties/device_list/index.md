---
page_title: "device_list"
subcategory: ""
description: "Add device for all interfaces belonging to this fleet."
xcsh_docs: {"aliases": ["device list"], "body_bytes": 1255, "body_sha256": "sha256:f34457ad3c4db4a5a56aa4ebbd61f177c95a1ddd9d69531bb8f02469ec476d26", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/device_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list"], "schema_version": 1, "sections": [{"aliases": ["devices"], "anchor": "section", "description": "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["device_list", "devices"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add device for all interfaces belonging to this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add device for all interfaces belonging to this fleet.

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
device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/): complete subsection reference.

## Next pages

- [device_list.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
