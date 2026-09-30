---
page_title: "client_side_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1704, "body_sha256": "sha256:f07a278e8ac2579ae3aee9f311d99f8d8893b6138ed9b05fe46315f97518b1c6", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy", "path": "docs/guides/data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# client_side_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [client_side_defense](data-sources--http_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](data-sources--http_loadbalancer--properties--client_side_defense--policy.md)
- client_side_defense.policy.js_insertion_rules

<a id="section"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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

- [exclude_list](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md): complete subsection reference.

- [rules](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md)
- [client_side_defense.policy](data-sources--http_loadbalancer--properties--client_side_defense--policy.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
