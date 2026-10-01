---
page_title: "rules.criteria.tcp.source_port"
subcategory: ""
description: "rules.criteria.tcp.source_port for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4070, "body_sha256": "sha256:71a44292b6859e8c06bb760b7b535f6615160d9bbee5477ffc856ce53ab64b77", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "path": "docs/guides/resources--nat_policy--properties--rules--criteria--tcp--source_port.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria", "tcp", "source_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/tcp/source_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria.tcp.source_port for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp.source_port

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- [rules.criteria](resources--nat_policy--properties--rules--criteria.md)
- [rules.criteria.tcp](resources--nat_policy--properties--rules--criteria--tcp.md)
- rules.criteria.tcp.source_port

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

- [no_port_match](resources--nat_policy--properties--rules--criteria--tcp--source_port--no_port_match.md): complete subsection reference.

<a id="schema-rules--criteria--tcp--source_port--port"></a>

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

<a id="schema-rules--criteria--tcp--source_port--port_ranges"></a>

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

- [rules.criteria.tcp.source_port.no_port_match](resources--nat_policy--properties--rules--criteria--tcp--source_port--no_port_match.md)
- [rules.criteria.tcp](resources--nat_policy--properties--rules--criteria--tcp.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
