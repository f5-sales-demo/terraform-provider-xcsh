---
page_title: "rules.spec.segment_policy.intra_segment"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules spec segment policy intra segment"], "body_bytes": 1628, "body_sha256": "sha256:9c2fb29aad359a08cb06652705bc23842f4fcfa608ed1c3895d39683876799df", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "path": "documentation/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1031111222233131-3030031320222313-1011000131211231-0212002101013303-2113213210020111-3011323310000203-0133311330023233-2212132203310000", "registry_path": "docs/guides/resources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "segment_policy", "intra_segment"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy.intra_segment

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
- [rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/)
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

- [rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
