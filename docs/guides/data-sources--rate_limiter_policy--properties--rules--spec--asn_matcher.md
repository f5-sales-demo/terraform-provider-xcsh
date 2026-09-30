---
page_title: "rules.spec.asn_matcher"
subcategory: "Security"
description: "rules.spec.asn_matcher for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1078, "body_sha256": "sha256:05febc081578149fafc7556806e2b4114bbb05fe43615f4cbdceaf94affb1f90", "canonical_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/data-sources--rate_limiter_policy--properties--rules--spec--asn_matcher.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.asn_matcher for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.spec.asn_matcher

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
- [Property reference](data-sources--rate_limiter_policy--reference.md)
- [rules](data-sources--rate_limiter_policy--properties--rules.md)
- [rules.spec](data-sources--rate_limiter_policy--properties--rules--spec.md)
- rules.spec.asn_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

- [asn_sets](data-sources--rate_limiter_policy--properties--rules--spec--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [rules.spec.asn_matcher.asn_sets](data-sources--rate_limiter_policy--properties--rules--spec--asn_matcher--asn_sets.md)
- [rules.spec](data-sources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
