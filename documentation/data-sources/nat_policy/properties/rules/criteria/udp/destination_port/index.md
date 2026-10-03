---
page_title: "rules.criteria.udp.destination_port"
subcategory: ""
description: "Port match of the request can be a range or a specific port."
xcsh_docs: {"aliases": ["rules criteria udp destination port"], "body_bytes": 3869, "body_sha256": "sha256:b0861dc873fffe82b69504816462fb6f21f5abf498b3b41250f4b4fd4aa8d313", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port:no_port_match"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp", "path": "documentation/data-sources/nat_policy/properties/rules/criteria/udp/destination_port/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2322103312012322-1302132112232110-1110123020031021-0220032131210330-3111220200222033-0101310201132011-0032023233130101-2202030332003020", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "udp", "destination_port"], "schema_version": 1, "sections": [{"aliases": ["rules criteria udp destination port no port match"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port:no_port_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "udp", "destination_port", "no_port_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria udp destination port port"], "anchor": "schema-rules--criteria--udp--destination_port--port", "description": "Exclusive with Exact Port to match.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "udp", "destination_port", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["rules criteria udp destination port port ranges"], "anchor": "schema-rules--criteria--udp--destination_port--port_ranges", "description": "Exclusive with Port range to match.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "udp", "destination_port", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/udp/destination_port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Port match of the request can be a range or a specific port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.udp.destination_port

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/)
- [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/udp/)
- rules.criteria.udp.destination_port

<a id="section"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

## Direct properties

- [no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/udp/destination_port/no_port_match/): complete subsection reference.

<a id="schema-rules--criteria--udp--destination_port--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-rules--criteria--udp--destination_port--port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

## Next pages

- [rules.criteria.udp.destination_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/udp/destination_port/no_port_match/)
- [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/udp/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
