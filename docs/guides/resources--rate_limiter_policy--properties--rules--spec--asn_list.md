---
page_title: "rules.spec.asn_list"
subcategory: "Security"
description: "rules.spec.asn_list for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3156, "body_sha256": "sha256:fbd3c113e5bd90495ac584ecd72d7734afdf9d9cc829955c772eef42acc4a5e1", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_list", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_list", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--asn_list.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "asn_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/asn_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.asn_list for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.spec.asn_list

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- rules.spec.asn_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--spec--asn_list--as_numbers"></a>

### as_numbers property

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
