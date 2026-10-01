---
page_title: "interface_list"
subcategory: ""
description: "interface_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1115, "body_sha256": "sha256:a0bca5c0ce701f87f7c91e44dbba7af9c93202ca7d46d42105fb1e2ea6a6e675", "canonical_id": "xcsh-docs:resources:fleet:properties:interface_list", "child_ids": ["xcsh-docs:resources:fleet:properties:interface_list:interfaces"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:interface_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--interface_list.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "interface_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# interface_list

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](resources--fleet--properties--interface_list--interfaces.md): complete subsection reference.

## Next pages

- [interface_list.interfaces](resources--fleet--properties--interface_list--interfaces.md)
- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
