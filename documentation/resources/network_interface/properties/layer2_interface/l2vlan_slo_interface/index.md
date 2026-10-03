---
page_title: "layer2_interface.l2vlan_slo_interface"
subcategory: ""
description: "Layer2 Site Local Outside VLAN Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface l2vlan slo interface"], "body_bytes": 2538, "body_sha256": "sha256:f11f2400f70b10da98031f7da0fe9aaf35ad32bfed6a82fe18fbd85d3f808b8b", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "parent_id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "path": "documentation/resources/network_interface/properties/layer2_interface/l2vlan_slo_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1302323103122132-1203113111202003-2133211110002320-1232221120230211-3000113113333120-0301230011130031-3323002003223012-2333231221102112", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "schema-layer2_interface--l2vlan_slo_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_slo_interface:RequiredObjectAttributes:vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface", "l2vlan_slo_interface"], "schema_version": 1, "sections": [{"aliases": ["layer2 interface l2vlan slo interface vlan id"], "anchor": "schema-layer2_interface--l2vlan_slo_interface--vlan_id", "description": "VLAN ID", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2vlan_slo_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/layer2_interface/l2vlan_slo_interface/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Layer2 Site Local Outside VLAN Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface.l2vlan_slo_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/)
- layer2_interface.l2vlan_slo_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Layer2 Site Local Outside VLAN Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vlan_id")}
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
l2vlan_slo_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-layer2_interface--l2vlan_slo_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

## Next pages

- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
