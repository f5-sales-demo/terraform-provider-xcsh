---
page_title: "rules.criteria.udp.source_port"
subcategory: ""
description: "Port match of the request can be a range or a specific port."
xcsh_docs: {"aliases": ["rules criteria udp source port"], "body_bytes": 4518, "body_sha256": "sha256:e142ec7642d4ea35a128c9c3f66f9e6b1bd5835cbf73d074c2a30441a0bc06c2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port:no_port_match"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "path": "documentation/resources/nat_policy/properties/rules/criteria/udp/source_port/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rules--criteria--udp--source_port--port", "enforcement": "provider-schema", "group": "rules.criteria.udp.source_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--udp--source_port--port", "enforcement": "provider-schema", "group": "rules.criteria.udp.source_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--udp--source_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.udp.source_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--udp--source_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.udp.source_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.udp.source_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.udp.source_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port:no_port_match", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "udp", "source_port"], "schema_version": 1, "sections": [{"aliases": ["rules criteria udp source port no port match"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port:no_port_match", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "udp", "source_port", "no_port_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria udp source port port"], "anchor": "schema-rules--criteria--udp--source_port--port", "description": "Exclusive with Exact Port to match.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "udp", "source_port", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["rules criteria udp source port port ranges"], "anchor": "schema-rules--criteria--udp--source_port--port_ranges", "description": "Exclusive with Port range to match.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp:source_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "udp", "source_port", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/udp/source_port/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Port match of the request can be a range or a specific port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.udp.source_port

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/)
- [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/)
- rules.criteria.udp.source_port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/no_port_match/): complete subsection reference.

<a id="schema-rules--criteria--udp--source_port--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-rules--criteria--udp--source_port--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [rules.criteria.udp.source_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/no_port_match/)
- [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
