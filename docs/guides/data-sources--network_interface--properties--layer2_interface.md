---
page_title: "layer2_interface"
subcategory: ""
description: "layer2_interface for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1709, "body_sha256": "sha256:7a4c09bbc54a23b6e6e5c5f9ce2c8709dcc2eefb439dfa434b89497a68771516", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2sriov_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "parent_id": "xcsh-docs:data-sources:network_interface:reference", "path": "docs/guides/data-sources--network_interface--properties--layer2_interface.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["layer2_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/layer2_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "layer2_interface for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- layer2_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for layer2 interface.

Upstream description:

Layer2 Interface Configuration.

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

## Direct properties

- [l2sriov_interface](data-sources--network_interface--properties--layer2_interface--l2sriov_interface.md): complete subsection reference.

- [l2vlan_interface](data-sources--network_interface--properties--layer2_interface--l2vlan_interface.md): complete subsection reference.

- [l2vlan_slo_interface](data-sources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md): complete subsection reference.

## Next pages

- [layer2_interface.l2sriov_interface](data-sources--network_interface--properties--layer2_interface--l2sriov_interface.md)
- [layer2_interface.l2vlan_interface](data-sources--network_interface--properties--layer2_interface--l2vlan_interface.md)
- [layer2_interface.l2vlan_slo_interface](data-sources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
