---
page_title: "storage_interface_list"
subcategory: ""
description: "storage_interface_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 901, "body_sha256": "sha256:d94cab4f88a31d19051b0418f0a505aa7901a5f0240cc5f401a71c7feaa99bcc", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_interface_list", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_interface_list:interfaces"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_interface_list", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "docs/guides/data-sources--fleet--properties--storage_interface_list.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_interface_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_interface_list

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- storage_interface_list

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

- [interfaces](data-sources--fleet--properties--storage_interface_list--interfaces.md): complete subsection reference.

## Next pages

- [storage_interface_list.interfaces](data-sources--fleet--properties--storage_interface_list--interfaces.md)
- [Property reference](data-sources--fleet--reference.md)
- [xcsh_fleet](../data-sources/fleet.md)
