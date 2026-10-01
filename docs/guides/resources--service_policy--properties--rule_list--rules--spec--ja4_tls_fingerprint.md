---
page_title: "rule_list.rules.spec.ja4_tls_fingerprint"
subcategory: "Security"
description: "rule_list.rules.spec.ja4_tls_fingerprint for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2798, "body_sha256": "sha256:33448c62075aa95eeccf5fe726686e345e8965d14a07269ba497fc49abdbcdf4", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--ja4_tls_fingerprint.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "ja4_tls_fingerprint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.ja4_tls_fingerprint for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.ja4_tls_fingerprint

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.ja4_tls_fingerprint

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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
ja4_tls_fingerprint {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--ja4_tls_fingerprint--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Optional.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../resources/service_policy.md)
