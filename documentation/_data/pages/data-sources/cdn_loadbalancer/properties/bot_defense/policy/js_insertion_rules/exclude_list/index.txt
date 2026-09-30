---
page_title: "bot_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "bot_defense.policy.js_insertion_rules.exclude_list for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3712, "body_sha256": "sha256:4e8d4e3cd851454fe84ade5df479ee50df21eeb3ba3d949923ebc6556c6726d1", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list:path"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bot_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.js_insertion_rules.exclude_list for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="section"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/path/): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/any_domain/)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/domain/)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/metadata/)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/exclude_list/path/)
- [bot_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
