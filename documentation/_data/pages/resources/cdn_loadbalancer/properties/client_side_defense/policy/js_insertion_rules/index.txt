---
page_title: "client_side_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "This defines custom JavaScript insertion rules for Client-Side Defense Policy."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules"], "body_bytes": 2530, "body_sha256": "sha256:8bceaeb934a044f236bebd9af4d98e48d769508a65f6164f2498486f22bd8025", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy", "path": "documentation/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "syntax": "block", "type": "object"}, {"aliases": ["rules"], "anchor": "section", "description": "Required list of pages to insert Client-Side Defense client JavaScript.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines custom JavaScript insertion rules for Client-Side Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/)
- client_side_defense.policy.js_insertion_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/): complete subsection reference.

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/)
- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
