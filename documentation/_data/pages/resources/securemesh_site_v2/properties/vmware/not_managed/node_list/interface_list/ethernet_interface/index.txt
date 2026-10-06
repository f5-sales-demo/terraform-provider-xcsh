---
page_title: "vmware.not_managed.node_list.interface_list.ethernet_interface"
subcategory: ""
description: "Configuration parameter for ethernet interface."
xcsh_docs: {"aliases": ["vmware not managed node list interface list ethernet interface"], "body_bytes": 4050, "body_sha256": "sha256:7319597647ed06829017de94f716d1255ae2b50ce65058ddbd10e3c4d6698a4e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ethernet_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-018.md", "relationships": [{"anchor": "schema-vmware--not_managed--node_list--interface_list--ethernet_interface--mac", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.ethernet_interface:RequiredObjectAttributes:mac", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ethernet_interface"], "schema_version": 1, "sections": [{"aliases": ["vmware not managed node list interface list ethernet interface device"], "anchor": "schema-vmware--not_managed--node_list--interface_list--ethernet_interface--device", "description": "Select an Ethernet device from the discovered interfaces to configure. Once configured, this interface will be part of this sites dataplane and can participate in the networking services configured on this site.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ethernet_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["vmware not managed node list interface list ethernet interface mac"], "anchor": "schema-vmware--not_managed--node_list--interface_list--ethernet_interface--mac", "description": "Configuration parameter for mac", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ethernet_interface", "mac"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ethernet_interface/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for ethernet interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list.interface_list.ethernet_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [vmware](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/)
- [vmware.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/)
- [vmware.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/)
- [vmware.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-vmware--not_managed--node_list--interface_list--ethernet_interface--device"></a>

### device property

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-vmware--not_managed--node_list--interface_list--ethernet_interface--mac"></a>

### mac property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```
