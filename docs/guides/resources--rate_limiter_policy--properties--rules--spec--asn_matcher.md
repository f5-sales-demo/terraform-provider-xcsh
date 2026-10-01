---
page_title: "rules.spec.asn_matcher"
subcategory: "Security"
description: "rules.spec.asn_matcher for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1420, "body_sha256": "sha256:44eed105752739b5da5cf222c7d13bbe3fe21ed87c04909e189a792d084d5fe2", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_matcher", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_matcher", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--asn_matcher.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.asn_matcher for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.asn_matcher

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- rules.spec.asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](resources--rate_limiter_policy--properties--rules--spec--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [rules.spec.asn_matcher.asn_sets](resources--rate_limiter_policy--properties--rules--spec--asn_matcher--asn_sets.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
