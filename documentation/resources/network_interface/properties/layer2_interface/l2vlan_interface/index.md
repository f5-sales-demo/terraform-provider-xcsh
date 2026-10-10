---
page_title: "layer2_interface.l2vlan_interface"
subcategory: ""
description: "Layer2 VLAN Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface l2vlan interface"], "body_bytes": 3545, "body_sha256": "sha256:b8c65d6b78fa21c5976e232dd7a78cc5d6611850a432829a2060e5db9e4fbc73", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "parent_id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "path": "documentation/resources/network_interface/properties/layer2_interface/l2vlan_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3020100313222301-3322210111012110-0002233303010223-2101112211013101-3110320121222020-1000123022030111-2201113203323033-3203202233203111", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "schema-layer2_interface--l2vlan_interface--device", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "type": "requires"}, {"anchor": "schema-layer2_interface--l2vlan_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface", "l2vlan_interface"], "schema_version": 1, "sections": [{"aliases": ["layer2 interface l2vlan interface device"], "anchor": "schema-layer2_interface--l2vlan_interface--device", "description": "Physical ethernet interface.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2vlan_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["layer2 interface l2vlan interface vlan id"], "anchor": "schema-layer2_interface--l2vlan_interface--vlan_id", "description": "VLAN ID", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2vlan_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/layer2_interface/l2vlan_interface/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Layer2 VLAN Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface.l2vlan_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/)
- layer2_interface.l2vlan_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for l2vlan interface.

Additional upstream details:

Layer2 VLAN Interface Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
l2vlan_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-layer2_interface--l2vlan_interface--device"></a>

### device property

Type: `"string"`. Optional.

Ethernet Device. Physical ethernet interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-layer2_interface--l2vlan_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

VLAN ID. VLAN ID

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
