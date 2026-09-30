---
page_title: "sriov_interfaces"
subcategory: ""
description: "sriov_interfaces for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 907, "body_sha256": "sha256:d9f01f1590bdb624683e659f99d883bd97c5328086fe5886089006d6879c5668", "canonical_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces", "child_ids": ["xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:sriov_interfaces", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--sriov_interfaces.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sriov_interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/sriov_interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sriov_interfaces for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# sriov_interfaces

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- sriov_interfaces

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

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
sriov_interfaces {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sriov_interface](resources--fleet--properties--sriov_interfaces--sriov_interface.md): complete subsection reference.

## Next pages

- [sriov_interfaces.sriov_interface](resources--fleet--properties--sriov_interfaces--sriov_interface.md)
- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
