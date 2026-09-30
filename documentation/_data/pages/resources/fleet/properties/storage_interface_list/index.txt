---
page_title: "storage_interface_list"
subcategory: ""
description: "storage_interface_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1372, "body_sha256": "sha256:fcea0fac2821eca388ed448cd670b155bd98a17942d99dabd805e556888f9b00", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_interface_list:interfaces"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_interface_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/storage_interface_list/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["storage_interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_interface_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_interface_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- storage_interface_list

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
storage_interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_interface_list/interfaces/): complete subsection reference.

## Next pages

- [storage_interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_interface_list/interfaces/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
