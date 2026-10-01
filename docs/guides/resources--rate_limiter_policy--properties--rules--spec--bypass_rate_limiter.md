---
page_title: "rules.spec.bypass_rate_limiter"
subcategory: "Security"
description: "rules.spec.bypass_rate_limiter for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1131, "body_sha256": "sha256:d21a9787deb002dd44639c581db9d82f1585b6afdf1791540d65a0a567232bdf", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--bypass_rate_limiter.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "bypass_rate_limiter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.bypass_rate_limiter for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.bypass_rate_limiter

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- rules.spec.bypass_rate_limiter

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for bypass rate limiter.

Upstream description:

This can be used for messages where no values are needed.

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
bypass_rate_limiter = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
