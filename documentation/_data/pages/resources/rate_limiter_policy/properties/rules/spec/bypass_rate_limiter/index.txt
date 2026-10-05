---
page_title: "rules.spec.bypass_rate_limiter"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules spec bypass rate limiter"], "body_bytes": 1437, "body_sha256": "sha256:6631c4f9130058c7e5a2fa149e12ab899ebe1b8f24a5abebd3e3b9ab8b3b3657", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "documentation/resources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3022322213330233-0231202100023331-3200222121011103-2033231011132032-3233222020200330-1310313113302210-2333012103310303-1030113210010233", "registry_path": "docs/guides/resources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "bypass_rate_limiter"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.bypass_rate_limiter

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
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

- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
