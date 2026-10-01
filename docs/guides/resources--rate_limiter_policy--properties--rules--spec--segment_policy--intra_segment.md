---
page_title: "rules.spec.segment_policy.intra_segment"
subcategory: "Security"
description: "rules.spec.segment_policy.intra_segment for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1274, "body_sha256": "sha256:86671b0252190b1e60adb83d66e4849bbcddbf205c920656a81d90b8f20897ad", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--segment_policy--intra_segment.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "segment_policy", "intra_segment"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.segment_policy.intra_segment for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy.intra_segment

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [rules.spec.segment_policy](resources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- rules.spec.segment_policy.intra_segment

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for intra segment.

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
intra_segment = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.spec.segment_policy](resources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
