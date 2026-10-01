---
page_title: "client_side_defense.policy.js_insertion_rules.rules"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules.rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3398, "body_sha256": "sha256:4b2eb2627dc9499b483b89ec82c15c6497b4a2f94b10b3ba8f56e4d3a990508a", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules.rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [client_side_defense](data-sources--cdn_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--properties--client_side_defense--policy.md)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- client_side_defense.policy.js_insertion_rules.rules

<a id="section"></a>

Type: `"list"`. Computed.

Required list of pages to insert Client-Side Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [any_domain](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--any_domain.md): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--domain.md): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--metadata.md): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--path.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--any_domain.md)
- [client_side_defense.policy.js_insertion_rules.rules.domain](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--domain.md)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--metadata.md)
- [client_side_defense.policy.js_insertion_rules.rules.path](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--path.md)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
