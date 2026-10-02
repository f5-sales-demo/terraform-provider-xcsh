---
page_title: "equinix.not_managed.node_list.interface_list.vlan_interface"
subcategory: ""
description: "Configuration parameter for vlan interface."
xcsh_docs: {"aliases": ["equinix not managed node list interface list vlan interface"], "body_bytes": 4220, "body_sha256": "sha256:4556b23bcd09fbfec2baa9f13f2443f09669900721cd2a8febad7221a7afd01e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/vlan_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1320201032311302-2010322011120020-2300202030302120-2310032210331223-3102333123002122-0011212110103133-2031023020012131-3013110312323121", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-009.md", "relationships": [{"anchor": "schema-equinix--not_managed--node_list--interface_list--vlan_interface--device", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list.vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "type": "requires"}, {"anchor": "schema-equinix--not_managed--node_list--interface_list--vlan_interface--vlan_id", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list.vlan_interface:RequiredObjectAttributes:device,vlan_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "vlan_interface"], "schema_version": 1, "sections": [{"aliases": ["device"], "anchor": "schema-equinix--not_managed--node_list--interface_list--vlan_interface--device", "description": "Select a parent interface from the dropdown.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "vlan_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["vlan id"], "anchor": "schema-equinix--not_managed--node_list--interface_list--vlan_interface--vlan_id", "description": "Configure the VLAN tag for this interface.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "vlan_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/vlan_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for vlan interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# equinix.not_managed.node_list.interface_list.vlan_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [equinix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/)
- [equinix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/)
- [equinix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/)
- [equinix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/)
- equinix.not_managed.node_list.interface_list.vlan_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
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

<a id="schema-equinix--not_managed--node_list--interface_list--vlan_interface--device"></a>

### device property

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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

<a id="schema-equinix--not_managed--node_list--interface_list--vlan_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

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

- [equinix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
