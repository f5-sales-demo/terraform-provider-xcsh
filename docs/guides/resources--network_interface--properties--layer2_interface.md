---
page_title: "layer2_interface"
subcategory: ""
description: "layer2_interface for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 2168, "body_sha256": "sha256:ce18c9b8f50a25f5b9ec9e8edfb163bc6e83e48d781c5b8f3b1a3bd62b724ef3", "canonical_id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "child_ids": ["xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "parent_id": "xcsh-docs:resources:network_interface:reference", "path": "docs/guides/resources--network_interface--properties--layer2_interface.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["layer2_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/layer2_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "layer2_interface for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- layer2_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for layer2 interface.

Upstream description:

Layer2 Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("l2sriov_interface",
    "l2vlan_interface"),
  validators.ConflictingObjectAttributes("l2sriov_interface",
    "l2vlan_slo_interface"),
  validators.ConflictingObjectAttributes("l2vlan_interface",
    "l2vlan_slo_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-layer2_interface_choice": "[\"l2sriov_interface\",\"l2vlan_interface\",\"l2vlan_slo_interface\"]"
}
```

Terraform syntax:

```terraform
layer2_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [l2sriov_interface](resources--network_interface--properties--layer2_interface--l2sriov_interface.md): complete subsection reference.

- [l2vlan_interface](resources--network_interface--properties--layer2_interface--l2vlan_interface.md): complete subsection reference.

- [l2vlan_slo_interface](resources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md): complete subsection reference.

## Next pages

- [layer2_interface.l2sriov_interface](resources--network_interface--properties--layer2_interface--l2sriov_interface.md)
- [layer2_interface.l2vlan_interface](resources--network_interface--properties--layer2_interface--l2vlan_interface.md)
- [layer2_interface.l2vlan_slo_interface](resources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [xcsh_network_interface](../resources/network_interface.md)
