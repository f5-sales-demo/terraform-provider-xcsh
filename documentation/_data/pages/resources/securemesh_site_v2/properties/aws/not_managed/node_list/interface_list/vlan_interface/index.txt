---
page_title: "aws.not_managed.node_list.interface_list.vlan_interface"
subcategory: ""
description: "Configuration parameter for vlan interface."
xcsh_docs: {"aliases": ["aws not managed node list interface list vlan interface"], "body_bytes": 3938, "body_sha256": "sha256:61ec0e71a5a8b7eb1b0f003d30cea4b6778c5fd1e53965051d0e47fa97b2a3e2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:vlan_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/vlan_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1021322130033210-3013010303233300-1230322322220200-1030112032003212-0101211221120122-0231232330311100-1012100021331010-1131310012212000", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-004.md", "relationships": [{"anchor": "schema-aws--not_managed--node_list--interface_list--vlan_interface--device", "enforcement": "provider-schema", "group": "aws.not_managed.node_list.interface_list.vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:vlan_interface", "type": "requires"}, {"anchor": "schema-aws--not_managed--node_list--interface_list--vlan_interface--vlan_id", "enforcement": "provider-schema", "group": "aws.not_managed.node_list.interface_list.vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:vlan_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "vlan_interface"], "schema_version": 1, "sections": [{"aliases": ["aws not managed node list interface list vlan interface device"], "anchor": "schema-aws--not_managed--node_list--interface_list--vlan_interface--device", "description": "Select a parent interface from the dropdown.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:vlan_interface", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "not_managed", "node_list", "interface_list", "vlan_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws not managed node list interface list vlan interface vlan id"], "anchor": "schema-aws--not_managed--node_list--interface_list--vlan_interface--vlan_id", "description": "Configure the VLAN tag for this interface.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "not_managed", "node_list", "interface_list", "vlan_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/vlan_interface/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration parameter for vlan interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.vlan_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/)
- [aws.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/)
- [aws.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/)
- [aws.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/)
- aws.not_managed.node_list.interface_list.vlan_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

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
vlan_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws--not_managed--node_list--interface_list--vlan_interface--device"></a>

### device property

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-aws--not_managed--node_list--interface_list--vlan_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```
