---
page_title: "interface_list"
subcategory: ""
description: "interface_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 762, "body_sha256": "sha256:4c0281cf575e006b45fd25cdbc5255f420fd4716819bfbbf1b7c34ab11c7552c", "canonical_id": "xcsh-docs:data-sources:fleet:properties:interface_list", "child_ids": ["xcsh-docs:data-sources:fleet:properties:interface_list:interfaces"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:interface_list", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "docs/guides/data-sources--fleet--properties--interface_list.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "interface_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# interface_list

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- interface_list

<a id="section"></a>

Type: `"single"`. Computed.

Add all interfaces belonging to this fleet.

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

- [interfaces](data-sources--fleet--properties--interface_list--interfaces.md): complete subsection reference.

## Next pages

- [interface_list.interfaces](data-sources--fleet--properties--interface_list--interfaces.md)
- [Property reference](data-sources--fleet--reference.md)
- [xcsh_fleet](../data-sources/fleet.md)
