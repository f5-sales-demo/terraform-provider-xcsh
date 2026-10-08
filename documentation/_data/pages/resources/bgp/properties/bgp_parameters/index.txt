---
page_title: "bgp_parameters"
subcategory: ""
description: "BGP parameters for the local site."
xcsh_docs: {"aliases": ["bgp parameters"], "body_bytes": 3786, "body_sha256": "sha256:b57bba4470b02e674b65bac3a63dc848a7dd514dd55d4590000806d4dfd820c9", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:bgp_parameters:from_site", "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "documentation/resources/bgp/properties/bgp_parameters/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3103003320212223-3323032330310220-0112102013012011-3303212132102001-2201310010330000-1120110231210311-2311000111033323-0133120330101211", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "schema-bgp_parameters--ip_address", "enforcement": "provider-schema", "group": "bgp_parameters:ConflictingObjectAttributes:from_site,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "type": "conflicts"}, {"anchor": "schema-bgp_parameters--ip_address", "enforcement": "provider-schema", "group": "bgp_parameters:ConflictingObjectAttributes:ip_address,local_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bgp_parameters:ConflictingObjectAttributes:from_site,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bgp_parameters:ConflictingObjectAttributes:from_site,local_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:from_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bgp_parameters:ConflictingObjectAttributes:from_site,local_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bgp_parameters:ConflictingObjectAttributes:ip_address,local_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address", "type": "conflicts"}, {"anchor": "schema-bgp_parameters--asn", "enforcement": "provider-schema", "group": "bgp_parameters:RequiredObjectAttributes:asn", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bgp_parameters"], "schema_version": 1, "sections": [{"aliases": ["bgp parameters asn"], "anchor": "schema-bgp_parameters--asn", "description": "Configure the autonomous system number (ASN) for the local site.", "document_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["bgp parameters from site"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:from_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "from_site"], "syntax": "attribute", "type": "object"}, {"aliases": ["bgp parameters ip address"], "anchor": "schema-bgp_parameters--ip_address", "description": "Exclusive with Use the configured IPv4 Address as Router ID.", "document_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["bgp parameters local address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "local_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/bgp_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "BGP parameters for the local site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bgpCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bgp_parameters

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- bgp_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bgp parameters.

Additional upstream details:

BGP parameters for the local site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn"),
  validators.ConflictingObjectAttributes("from_site",
    "ip_address"),
  validators.ConflictingObjectAttributes("from_site",
    "local_address"),
  validators.ConflictingObjectAttributes("ip_address",
    "local_address")}
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
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

Terraform syntax:

```terraform
bgp_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bgp_parameters--asn"></a>

### asn property

Type: `"number"`. Optional.

ASN. Autonomous System Number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/from_site/): complete subsection reference.

<a id="schema-bgp_parameters--ip_address"></a>

### ip_address property

Type: `"string"`. Optional.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [local_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/local_address/): complete subsection reference.
