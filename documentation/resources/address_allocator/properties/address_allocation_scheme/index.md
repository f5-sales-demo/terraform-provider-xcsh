---
page_title: "address_allocation_scheme"
subcategory: ""
description: "Decides the scheme to be used to allocate addresses from the configured address pool."
xcsh_docs: {"aliases": ["address allocation scheme"], "body_bytes": 7202, "body_sha256": "sha256:47e576f56ad0eaf0c95900020d285ac7b86da8f5df1e0d71d668b28396b6b39d", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:address_allocator:collection", "completeness": "complete", "id": "xcsh-docs:resources:address_allocator:properties:address_allocation_scheme", "parent_id": "xcsh-docs:resources:address_allocator:reference", "path": "documentation/resources/address_allocator/properties/address_allocation_scheme/index.md", "product": "distributed-cloud", "provider_name": "address_allocator", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3230300123333023-3021311210200313-1101220223020003-1031221103211013-2001322121201301-0211230330330103-2030331033011233-2322001313001032", "registry_path": "docs/guides/resources--address_allocator--reference--group-001.md", "relationships": [{"anchor": "schema-address_allocation_scheme--allocation_unit", "enforcement": "provider-schema", "group": "address_allocation_scheme:RequiredObjectAttributes:allocation_unit", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:address_allocator:properties:address_allocation_scheme", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["address_allocation_scheme"], "schema_version": 1, "sections": [{"aliases": ["address allocation scheme allocation unit"], "anchor": "schema-address_allocation_scheme--allocation_unit", "description": "Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30, subnets of /30 will be allocated from the given address pool.", "document_id": "xcsh-docs:resources:address_allocator:properties:address_allocation_scheme", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_allocation_scheme", "allocation_unit"], "syntax": "attribute", "type": "number"}, {"aliases": ["address allocation scheme local interface address offset"], "anchor": "schema-address_allocation_scheme--local_interface_address_offset", "description": "This is used to derive address for the local interface from the allocated subnet. If Local Interface Address Type is set to \"Offset from beginning of Subnet\", this offset value is added to the allocated subnet and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24 and offset is", "document_id": "xcsh-docs:resources:address_allocator:properties:address_allocation_scheme", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_allocation_scheme", "local_interface_address_offset"], "syntax": "attribute", "type": "number"}, {"aliases": ["address allocation scheme local interface address type"], "anchor": "schema-address_allocation_scheme--local_interface_address_type", "description": "Dictates how local interface address is derived from the allocated subnet Use Nth address of the allocated subnet as the local interface address, N being the Local Interface Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and Local Interface Address Type", "document_id": "xcsh-docs:resources:address_allocator:properties:address_allocation_scheme", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["LOCAL_INTERFACE_ADDRESS_FROM_PREFIX", "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN", "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_allocation_scheme", "local_interface_address_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/address_allocator/properties/address_allocation_scheme/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Decides the scheme to be used to allocate addresses from the configured address pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# address_allocation_scheme

Breadcrumbs:

- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/properties/)
- address_allocation_scheme

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Decides the scheme to be used to allocate addresses from the configured address pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("allocation_unit")}
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
address_allocation_scheme {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-address_allocation_scheme--allocation_unit"></a>

### allocation_unit property

Type: `"number"`. Optional.

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-address_allocation_scheme--local_interface_address_offset"></a>

### local_interface_address_offset property

Type: `"number"`. Optional.

Used to derive address for the local interface from the allocated subnet. If Local Interface Address
Type is set to 'Offset from beginning of Subnet', this offset value is added to the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24..

Additional upstream details:

This is used to derive address for the local interface from the allocated subnet. If Local Interface
Address Type is set to "Offset from beginning of Subnet", this offset value is added to the
allocated subnet and used as the local interface address. For example, if the allocated subnet is
192.0.2.0/24 and offset is set to 2 with Local Interface Address Type set to "Offset from beginning
of Subnet", local interface address of 192.0.2.204 is used. If Local Interface Address Type is set
to "Offset from end of Subnet", this offset value is subtracted from the end of the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24 and
offset is set to 1 with Local Interface Address Type set to "Offset from end of Subnet", local
interface address of 192.0.2.204 is used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-address_allocation_scheme--local_interface_address_type"></a>

### local_interface_address_type property

Type: `"string"`. Optional.

\[Enum:
LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN|LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END|LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\]
Dictates how local interface address is derived from the allocated subnet Use Nth address of the
allocated subnet as the local interface address, N being the Local Interface Address Offset. For
example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and
Local.. Possible values are \`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`,
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END\`,
\`LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\`. Defaults to
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`.

Additional upstream details:

Dictates how local interface address is derived from the allocated subnet

Use Nth address of the allocated subnet as the local interface address, N being the Local Interface
Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset
is set to 2 and Local Interface Address Type is set to "Offset from beginning of Subnet", local
address of 192.0.2.204 is used. Use Nth last address of the allocated subnet as the local interface
address, N being the Local Interface Address Offset. For example, if the allocated subnet is
192.0.2.0/24, Local Interface Address Offset is set to 1 and Local Interface Address Type is set to
"Offset from end of Subnet", local address of 192.0.2.204 is used. This case is used for
external\_connector.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LOCAL_INTERFACE_ADDRESS_FROM_PREFIX","LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN","LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
  "enum": [
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
