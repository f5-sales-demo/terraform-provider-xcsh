---
page_title: "rules.spec.segment_policy.dst_segments"
subcategory: "Security"
description: "rules.spec.segment_policy.dst_segments for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1436, "body_sha256": "sha256:42fbe1aa452eb8015da1b635144c71e3b5ebcd5e0be8c1b197c8dfe69b9990d1", "canonical_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments:segments"], "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy", "path": "docs/guides/data-sources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_segments.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "segment_policy", "dst_segments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.segment_policy.dst_segments for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy.dst_segments

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
- [Property reference](data-sources--rate_limiter_policy--reference.md)
- [rules](data-sources--rate_limiter_policy--properties--rules.md)
- [rules.spec](data-sources--rate_limiter_policy--properties--rules--spec.md)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- rules.spec.segment_policy.dst_segments

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

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

## Direct properties

- [segments](data-sources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_segments--segments.md): complete subsection reference.

## Next pages

- [rules.spec.segment_policy.dst_segments.segments](data-sources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_segments--segments.md)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
