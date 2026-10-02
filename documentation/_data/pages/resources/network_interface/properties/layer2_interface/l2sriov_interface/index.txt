---
page_title: "layer2_interface.l2sriov_interface"
subcategory: ""
description: "Layer2 SR-IOV Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface l2sriov interface"], "body_bytes": 4300, "body_sha256": "sha256:6c768586cc0272cfd040ce76564c6f386453f6e18c8a7e9db3465fe7b7753828", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface:untagged"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "parent_id": "xcsh-docs:resources:network_interface:properties:layer2_interface", "path": "documentation/resources/network_interface/properties/layer2_interface/l2sriov_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2001213330111322-1102020311010223-3113031102330111-3233331202011220-0021313020301221-1013000221232132-1311230023032233-0232100011312112", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "schema-layer2_interface--l2sriov_interface--vlan_id", "enforcement": "provider-schema", "group": "layer2_interface.l2sriov_interface:ConflictingObjectAttributes:untagged,vlan_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "layer2_interface.l2sriov_interface:ConflictingObjectAttributes:untagged,vlan_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface:untagged", "type": "conflicts"}, {"anchor": "schema-layer2_interface--l2sriov_interface--device", "enforcement": "provider-schema", "group": "layer2_interface.l2sriov_interface:RequiredObjectAttributes:device", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface", "l2sriov_interface"], "schema_version": 1, "sections": [{"aliases": ["device"], "anchor": "schema-layer2_interface--l2sriov_interface--device", "description": "Physical ethernet interface.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2sriov_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["untagged"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface:untagged", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2sriov_interface", "untagged"], "syntax": "attribute", "type": "object"}, {"aliases": ["vlan id"], "anchor": "schema-layer2_interface--l2sriov_interface--vlan_id", "description": "Exclusive with Configure a VLAN tagged interface.", "document_id": "xcsh-docs:resources:network_interface:properties:layer2_interface:l2sriov_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["layer2_interface", "l2sriov_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/layer2_interface/l2sriov_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Layer2 SR-IOV Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface.l2sriov_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/)
- layer2_interface.l2sriov_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for l2sriov interface.

Upstream description:

Layer2 SR-IOV Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("untagged",
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
  },
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
l2sriov_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-layer2_interface--l2sriov_interface--device"></a>

### device property

Type: `"string"`. Optional.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2sriov_interface/untagged/): complete subsection reference.

<a id="schema-layer2_interface--l2sriov_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged interface.

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
    "create": false,
    "minimum_config": false,
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

## Next pages

- [layer2_interface.l2sriov_interface.untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/l2sriov_interface/untagged/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
