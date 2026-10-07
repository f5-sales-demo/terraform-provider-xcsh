---
page_title: "bgp_parameters"
subcategory: ""
description: "BGP parameters for the local site."
xcsh_docs: {"aliases": ["bgp parameters"], "body_bytes": 2915, "body_sha256": "sha256:50714e2708ca13aa89d30450b4ea2d2cb82d2e973cc6478b79299915777306b7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:bgp_parameters:from_site", "xcsh-docs:data-sources:bgp:properties:bgp_parameters:local_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters", "parent_id": "xcsh-docs:data-sources:bgp:reference", "path": "documentation/data-sources/bgp/properties/bgp_parameters/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bgp_parameters"], "schema_version": 1, "sections": [{"aliases": ["bgp parameters asn"], "anchor": "schema-bgp_parameters--asn", "description": "Inspect the autonomous system number (ASN) for the local site.", "document_id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["bgp parameters from site"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters:from_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "from_site"], "syntax": "attribute", "type": "object"}, {"aliases": ["bgp parameters ip address"], "anchor": "schema-bgp_parameters--ip_address", "description": "Exclusive with Use the configured IPv4 Address as Router ID.", "document_id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["bgp parameters local address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters:local_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bgp_parameters", "local_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/bgp_parameters/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "BGP parameters for the local site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["bgpCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bgp_parameters

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- bgp_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for bgp parameters.

Additional upstream details:

BGP parameters for the local site.

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

## Direct properties

<a id="schema-bgp_parameters--asn"></a>

### asn property

Type: `"number"`. Computed.

ASN. Autonomous System Number.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/from_site/): complete subsection reference.

<a id="schema-bgp_parameters--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [local_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/local_address/): complete subsection reference.
