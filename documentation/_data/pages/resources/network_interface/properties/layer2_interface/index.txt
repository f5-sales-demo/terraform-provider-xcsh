---
page_title: "layer2_interface"
subcategory: ""
description: "Layer2 Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface"], "body_bytes": 2676, "body_sha256": "sha256:ba8bcf21039ca5fb15352a2a5d2303c6f3e8d203163d013f29150e3dc4df606b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "parent_id": "xcsh-docs:resources:network_interface:reference", "path": "documentation/resources/network_interface/properties/layer2_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0130112110231113-3130032011200330-3231321010032002-1110233231103021-2000032102333133-0312112112133231-3001223102103322-3003100130233100", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface:ConflictingObjectAttributes:l2sriov_interface,l2vlan_interface", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface:ConflictingObjectAttributes:l2sriov_interface,l2vlan_slo_interface", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface:ConflictingObjectAttributes:l2sriov_interface,l2vlan_interface", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface:ConflictingObjectAttributes:l2vlan_interface,l2vlan_slo_interface", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface:ConflictingObjectAttributes:l2sriov_interface,l2vlan_slo_interface", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface:ConflictingObjectAttributes:l2vlan_interface,l2vlan_slo_interface", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface"], "schema_version": 1, "sections": [{"aliases": ["l2sriov interface"], "anchor": "section", "description": "Layer2 SR-IOV Interface Configuration.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-layer2_interface--l2sriov_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2sriov_interface:ConflictingObjectAttributes:untagged,vlan_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface.l2sriov_interface:ConflictingObjectAttributes:untagged,vlan_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface:untagged", "type": "conflicts"}, {"anchor": "schema-layer2_interface--l2sriov_interface--device", "enforcement": "provider-schema", "group": "layer2_interface.l2sriov_interface:RequiredObjectAttributes:device", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "type": "requires"}], "schema_path": ["layer2_interface", "l2sriov_interface"], "syntax": "block", "type": "object"}, {"aliases": ["l2vlan interface"], "anchor": "section", "description": "Layer2 VLAN Interface Configuration.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-layer2_interface--l2vlan_interface--device", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "type": "requires"}, {"anchor": "schema-layer2_interface--l2vlan_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "type": "requires"}], "schema_path": ["layer2_interface", "l2vlan_interface"], "syntax": "block", "type": "object"}, {"aliases": ["l2vlan slo interface"], "anchor": "section", "description": "Layer2 Site Local Outside VLAN Interface Configuration.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-layer2_interface--l2vlan_slo_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_slo_interface:RequiredObjectAttributes:vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "type": "requires"}], "schema_path": ["layer2_interface", "l2vlan_slo_interface"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/layer2_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Layer2 Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
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

- [l2sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2sriov_interface/): complete subsection reference.

- [l2vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2vlan_interface/): complete subsection reference.

- [l2vlan_slo_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2vlan_slo_interface/): complete subsection reference.

## Next pages

- [layer2_interface.l2sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2sriov_interface/)
- [layer2_interface.l2vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2vlan_interface/)
- [layer2_interface.l2vlan_slo_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2vlan_slo_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
