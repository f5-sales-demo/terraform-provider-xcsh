---
page_title: "layer2_interface.l2vlan_interface"
subcategory: ""
description: "Layer2 VLAN Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface l2vlan interface"], "body_bytes": 2904, "body_sha256": "sha256:af7cccc1adff01393fccf2b159d4ddc9140e61ba5909b638749a839e49d2bebc", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "parent_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "path": "documentation/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0330330120020103-1222300012103030-1232232331201001-2303012223312000-1223302103300000-1221113313122123-1013010302300132-0102113300031232", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface", "l2vlan_interface"], "schema_version": 1, "sections": [{"aliases": ["layer2 interface l2vlan interface device"], "anchor": "schema-layer2_interface--l2vlan_interface--device", "description": "Physical ethernet interface.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2vlan_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["layer2 interface l2vlan interface vlan id"], "anchor": "schema-layer2_interface--l2vlan_interface--vlan_id", "description": "VLAN ID", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2vlan_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Layer2 VLAN Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface.l2vlan_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/)
- layer2_interface.l2vlan_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for l2vlan interface.

Additional upstream details:

Layer2 VLAN Interface Configuration.

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

<a id="schema-layer2_interface--l2vlan_interface--device"></a>

### device property

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

Type: `"number"`. Computed.

VLAN ID. VLAN ID

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
