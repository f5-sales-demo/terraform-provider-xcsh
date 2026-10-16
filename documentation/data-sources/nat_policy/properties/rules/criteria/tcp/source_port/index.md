---
page_title: "rules.criteria.tcp.source_port"
subcategory: ""
description: "Port match of the request can be a range or a specific port."
xcsh_docs: {"aliases": ["rules criteria tcp source port"], "body_bytes": 3216, "body_sha256": "sha256:91e5fbbc88d8cdc3af1cd9bdf66c2d142ca5f0b45346e90adb9379d04ce4d8a5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "path": "documentation/data-sources/nat_policy/properties/rules/criteria/tcp/source_port/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0120022133030120-1101233122023312-3011033131103313-3302323212102310-0231000100120012-2002033230300210-1023333131311030-0010112120321101", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "tcp", "source_port"], "schema_version": 1, "sections": [{"aliases": ["rules criteria tcp source port no port match"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "tcp", "source_port", "no_port_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria tcp source port port"], "anchor": "schema-rules--criteria--tcp--source_port--port", "description": "Exclusive with Exact Port to match.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "tcp", "source_port", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["rules criteria tcp source port port ranges"], "anchor": "schema-rules--criteria--tcp--source_port--port_ranges", "description": "Exclusive with Port range to match.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "tcp", "source_port", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/tcp/source_port/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Port match of the request can be a range or a specific port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["nat_policyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp.source_port

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/)
- [rules.criteria.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/)
- rules.criteria.tcp.source_port

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

- [no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/source_port/no_port_match/): complete subsection reference.

<a id="schema-rules--criteria--tcp--source_port--port"></a>

### port property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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

<a id="schema-rules--criteria--tcp--source_port--port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
