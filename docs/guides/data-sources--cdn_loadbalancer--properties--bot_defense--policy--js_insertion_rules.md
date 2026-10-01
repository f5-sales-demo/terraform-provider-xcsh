---
page_title: "bot_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "bot_defense.policy.js_insertion_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1663, "body_sha256": "sha256:0d8cadc45da0aac842f1f06fc2ff81a00f50efb89b89f9588aa43cdf90c1860d", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "js_insertion_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.js_insertion_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [bot_defense](data-sources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--cdn_loadbalancer--properties--bot_defense--policy.md)
- bot_defense.policy.js_insertion_rules

<a id="section"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

- [exclude_list](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--exclude_list.md): complete subsection reference.

- [rules](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules.md): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--exclude_list.md)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules.md)
- [bot_defense.policy](data-sources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
