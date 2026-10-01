---
page_title: "rules.spec.any_country"
subcategory: "Security"
description: "rules.spec.any_country for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1099, "body_sha256": "sha256:6f34a674ae6371388b68f3bf46513e3a25add1970dfb4658071367114340f257", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_country", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_country", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--any_country.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "any_country"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/any_country/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.any_country for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.any_country

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- rules.spec.any_country

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for any country.

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
any_country = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
