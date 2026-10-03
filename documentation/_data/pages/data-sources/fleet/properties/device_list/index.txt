---
page_title: "device_list"
subcategory: ""
description: "Add device for all interfaces belonging to this fleet."
xcsh_docs: {"aliases": ["device list"], "body_bytes": 1153, "body_sha256": "sha256:161aafc494f006d12957aee457415de6bfb4b0d62487489c4c32180d6522d12a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:device_list:devices"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:device_list", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/device_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0322000310332200-0133003311233211-2200301200221003-0000110213122330-0302212202302221-1101312120312330-2230321020131013-1021211013301323", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list"], "schema_version": 1, "sections": [{"aliases": ["device list devices"], "anchor": "section", "description": "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node", "document_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["device_list", "devices"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/device_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add device for all interfaces belonging to this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- device_list

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/): complete subsection reference.

## Next pages

- [device_list.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
