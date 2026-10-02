---
page_title: "layer2_interface.l2vlan_slo_interface"
subcategory: ""
description: "Layer2 Site Local Outside VLAN Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface l2vlan slo interface"], "body_bytes": 2538, "body_sha256": "sha256:09b5f7172699a2f98e7f02d4dfe5f70d216fa09e7969c857d284406f89c3c3b0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "parent_id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "path": "documentation/resources/network_interface/properties/layer2_interface/l2vlan_slo_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1302323103122132-1203113111202003-2133211110002320-1232221120230211-3000113113333120-0301230011130031-3323002003223012-2333231221102112", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "schema-layer2_interface--l2vlan_slo_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_slo_interface:RequiredObjectAttributes:vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface", "l2vlan_slo_interface"], "schema_version": 1, "sections": [{"aliases": ["vlan id"], "anchor": "schema-layer2_interface--l2vlan_slo_interface--vlan_id", "description": "VLAN ID", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2vlan_slo_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/layer2_interface/l2vlan_slo_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Layer2 Site Local Outside VLAN Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
