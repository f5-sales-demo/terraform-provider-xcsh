---
page_title: "storage_static_routes.storage_routes.nexthop"
subcategory: ""
description: "Identifies the next-hop for a route."
xcsh_docs: {"aliases": ["storage static routes storage routes nexthop"], "body_bytes": 3334, "body_sha256": "sha256:8cc9fecd3bdd9232dfb8a227c149b946386a3ba8f8c3c52f74e5b40b06e7275f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:interface", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "path": "documentation/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop"], "schema_version": 1, "sections": [{"aliases": ["storage static routes storage routes nexthop interface"], "anchor": "section", "description": "Nexthop is network interface when type is \"Network-Interface\"", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:interface", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["storage static routes storage routes nexthop nexthop address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv6", "type": "conflicts"}], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address"], "syntax": "block", "type": "object"}, {"aliases": ["storage static routes storage routes nexthop type"], "anchor": "schema-storage_static_routes--storage_routes--nexthop--type", "description": "Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes there is only one local interface on the virtual network. Use the specified address as nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["NEXT_HOP_DEFAULT_GATEWAY", "NEXT_HOP_NETWORK_INTERFACE", "NEXT_HOP_USE_CONFIGURED"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Identifies the next-hop for a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fleetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.nexthop

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/)
- storage_static_routes.storage_routes.nexthop

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

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
nexthop {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/interface/): complete subsection reference.

- [nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/): complete subsection reference.

<a id="schema-storage_static_routes--storage_routes--nexthop--type"></a>

### type property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NEXT_HOP_DEFAULT_GATEWAY","NEXT_HOP_NETWORK_INTERFACE","NEXT_HOP_USE_CONFIGURED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
