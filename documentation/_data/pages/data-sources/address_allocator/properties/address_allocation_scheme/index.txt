---
page_title: "address_allocation_scheme"
subcategory: ""
description: "Decides the scheme to be used to allocate addresses from the configured address pool."
xcsh_docs: {"aliases": ["address allocation scheme"], "body_bytes": 6347, "body_sha256": "sha256:c00e9a40de4a47bb337b79312def333bc174f06623a79adab18fbe8e3ab279ce", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:address_allocator:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:address_allocator:properties:address_allocation_scheme", "parent_id": "xcsh-docs:data-sources:address_allocator:reference", "path": "documentation/data-sources/address_allocator/properties/address_allocation_scheme/index.md", "product": "distributed-cloud", "provider_name": "address_allocator", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2111321100330323-3002321212213203-0330313012110130-1321112130102233-1320122010220022-0000130220031132-2330210313130332-1223112100032321", "registry_path": "docs/guides/data-sources--address_allocator--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["address_allocation_scheme"], "schema_version": 1, "sections": [{"aliases": ["allocation unit"], "anchor": "schema-address_allocation_scheme--allocation_unit", "description": "Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30, subnets of /30 will be allocated from the given address pool.", "document_id": "xcsh-docs:data-sources:address_allocator:properties:address_allocation_scheme", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_allocation_scheme", "allocation_unit"], "syntax": "attribute", "type": "number"}, {"aliases": ["local interface address offset"], "anchor": "schema-address_allocation_scheme--local_interface_address_offset", "description": "This is used to derive address for the local interface from the allocated subnet. If Local Interface Address Type is set to \"Offset from beginning of Subnet\", this offset value is added to the allocated subnet and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24 and offset is", "document_id": "xcsh-docs:data-sources:address_allocator:properties:address_allocation_scheme", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_allocation_scheme", "local_interface_address_offset"], "syntax": "attribute", "type": "number"}, {"aliases": ["local interface address type"], "anchor": "schema-address_allocation_scheme--local_interface_address_type", "description": "Dictates how local interface address is derived from the allocated subnet Use Nth address of the allocated subnet as the local interface address, N being the Local Interface Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and Local Interface Address Type", "document_id": "xcsh-docs:data-sources:address_allocator:properties:address_allocation_scheme", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_allocation_scheme", "local_interface_address_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/address_allocator/properties/address_allocation_scheme/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Decides the scheme to be used to allocate addresses from the configured address pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# address_allocation_scheme

Breadcrumbs:

- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/)
- address_allocation_scheme

<a id="section"></a>

Type: `"single"`. Computed.

Decides the scheme to be used to allocate addresses from the configured address pool.

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

<a id="schema-address_allocation_scheme--allocation_unit"></a>

### allocation_unit property

Type: `"number"`. Computed.

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Upstream description:

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"number"`. Computed.

Used to derive address for the local interface from the allocated subnet. If Local Interface Address
Type is set to 'Offset from beginning of Subnet', this offset value is added to the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24..

Upstream description:

This is used to derive address for the local interface from the allocated subnet.

If Local Interface Address Type is set to "Offset from beginning of Subnet", this offset value is
added to the allocated subnet and used as the local interface address. For example, if the allocated
subnet is 192.0.2.0/24 and offset is set to 2 with Local Interface Address Type set to "Offset from
beginning of Subnet", local interface address of 192.0.2.204 is used.

If Local Interface Address Type is set to "Offset from end of Subnet", this offset value is
subtracted from the end of the allocated subnet and used as the local interface address. For
example, if the allocated subnet is 192.0.2.0/24 and offset is set to 1 with Local Interface Address
Type set to "Offset from end of Subnet", local interface address of 192.0.2.204 is used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Computed.

\[Enum:
LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN|LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END|LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\]
Dictates how local interface address is derived from the allocated subnet Use Nth address of the
allocated subnet as the local interface address, N being the Local Interface Address Offset. For
example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and
Local.. Possible values are \`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`,
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END\`,
\`LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\`. Defaults to
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`.

Upstream description:

Dictates how local interface address is derived from the allocated subnet

Use Nth address of the allocated subnet as the local interface address, N being the Local Interface
Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset
is set to 2 and Local Interface Address Type is set to "Offset from beginning of Subnet", local
address of 192.0.2.204 is used.

Use Nth last address of the allocated subnet as the local interface address, N being the Local
Interface Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface
Address Offset is set to 1 and Local Interface Address Type is set to "Offset from end of Subnet",
local address of 192.0.2.204 is used.

This case is used for external\_connector.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/)
- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/)
