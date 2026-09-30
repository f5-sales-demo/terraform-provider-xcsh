---
page_title: "rules.criteria.udp.destination_port"
subcategory: ""
description: "rules.criteria.udp.destination_port for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3322, "body_sha256": "sha256:ea94f62af7afdc5d5a4323b9fc30a41dcda41c87476f76e966dd284279e8c9c4", "canonical_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port:no_port_match"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp:destination_port", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:udp", "path": "docs/guides/data-sources--nat_policy--properties--rules--criteria--udp--destination_port.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria", "udp", "destination_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/udp/destination_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria.udp.destination_port for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.criteria.udp.destination_port

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md)
- [Property reference](data-sources--nat_policy--reference.md)
- [rules](data-sources--nat_policy--properties--rules.md)
- [rules.criteria](data-sources--nat_policy--properties--rules--criteria.md)
- [rules.criteria.udp](data-sources--nat_policy--properties--rules--criteria--udp.md)
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

- [no_port_match](data-sources--nat_policy--properties--rules--criteria--udp--destination_port--no_port_match.md): complete subsection reference.

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

- [rules.criteria.udp.destination_port.no_port_match](data-sources--nat_policy--properties--rules--criteria--udp--destination_port--no_port_match.md)
- [rules.criteria.udp](data-sources--nat_policy--properties--rules--criteria--udp.md)
- [xcsh_nat_policy](../data-sources/nat_policy.md)
