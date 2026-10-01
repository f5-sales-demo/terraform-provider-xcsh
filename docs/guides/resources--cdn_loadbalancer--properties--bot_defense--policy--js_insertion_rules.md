---
page_title: "bot_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "bot_defense.policy.js_insertion_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1904, "body_sha256": "sha256:e2a7ef9cee293143420151f0ad4c8feae18fd779ebf8c44e887a69511b363aa2", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules:rules"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "js_insertion_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.js_insertion_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- bot_defense.policy.js_insertion_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
```

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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_list](resources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--exclude_list.md): complete subsection reference.

- [rules](resources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules.md): complete subsection reference.

## Next pages

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--exclude_list.md)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules--rules.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
