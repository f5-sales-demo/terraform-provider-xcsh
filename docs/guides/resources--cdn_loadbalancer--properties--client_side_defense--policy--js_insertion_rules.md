---
page_title: "client_side_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2032, "body_sha256": "sha256:620b8eb5af0a51f7115667a477e89ab88cc5391ceb9b4ad5383b522f725f5a42", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy", "path": "docs/guides/resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [client_side_defense](resources--cdn_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](resources--cdn_loadbalancer--properties--client_side_defense--policy.md)
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

- [exclude_list](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md): complete subsection reference.

- [rules](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md)
- [client_side_defense.policy](resources--cdn_loadbalancer--properties--client_side_defense--policy.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
