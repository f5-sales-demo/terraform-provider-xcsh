---
page_title: "client_side_defense.policy.js_insertion_rules.exclude_list.any_domain"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules.exclude_list.any_domain for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1627, "body_sha256": "sha256:e6cdba2116759fc7030233e2a5065fa07a6730764ded391ed3ebb717389ddede", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "path": "docs/guides/resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--any_domain.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "any_domain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/any_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules.exclude_list.any_domain for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [client_side_defense](resources--http_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](resources--http_loadbalancer--properties--client_side_defense--policy.md)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
any_domain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
