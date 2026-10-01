---
page_title: "policy_based_challenge.rule_list.rules"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2733, "body_sha256": "sha256:696f6e5af671810620916f57791888ea08f1e196e2b1e1cca89dbe6a616a6146", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](data-sources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- policy_based_challenge.rule_list.rules

<a id="section"></a>

Type: `"list"`. Computed.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [metadata](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--metadata.md): complete subsection reference.

- [spec](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.metadata](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--metadata.md)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
