---
page_title: "rules.spec.ip_matcher"
subcategory: "Security"
description: "rules.spec.ip_matcher for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1688, "body_sha256": "sha256:f50505324e16191543ead4746636cc627e922bb98da4263207fc7e9829b4983d", "canonical_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher:prefix_sets"], "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/data-sources--rate_limiter_policy--properties--rules--spec--ip_matcher.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "ip_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.ip_matcher for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.spec.ip_matcher

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
- [Property reference](data-sources--rate_limiter_policy--reference.md)
- [rules](data-sources--rate_limiter_policy--properties--rules.md)
- [rules.spec](data-sources--rate_limiter_policy--properties--rules--spec.md)
- rules.spec.ip_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="schema-rules--spec--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--rate_limiter_policy--properties--rules--spec--ip_matcher--prefix_sets.md): complete subsection reference.

## Next pages

- [rules.spec.ip_matcher.prefix_sets](data-sources--rate_limiter_policy--properties--rules--spec--ip_matcher--prefix_sets.md)
- [rules.spec](data-sources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md)
