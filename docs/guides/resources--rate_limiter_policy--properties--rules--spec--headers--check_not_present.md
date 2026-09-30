---
page_title: "rules.spec.headers.check_not_present"
subcategory: "Security"
description: "rules.spec.headers.check_not_present for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1149, "body_sha256": "sha256:f45ac2889c1c507274362718a93d380e0d1b6ea47cc14730fb8ad35f5d622f4b", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:headers:check_not_present", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:headers:check_not_present", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:headers", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--headers--check_not_present.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "headers", "check_not_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/headers/check_not_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.headers.check_not_present for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.spec.headers.check_not_present

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [rules.spec.headers](resources--rate_limiter_policy--properties--rules--spec--headers.md)
- rules.spec.headers.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.spec.headers](resources--rate_limiter_policy--properties--rules--spec--headers.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
