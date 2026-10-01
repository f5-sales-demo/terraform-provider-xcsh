---
page_title: "client_side_defense.policy"
subcategory: "Load Balancing"
description: "client_side_defense.policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2371, "body_sha256": "sha256:956bcad829b08be13a28889d86943c2ff7ccee4ead9abcb7bb3bd3d90fa2d64d", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense", "path": "docs/guides/data-sources--http_loadbalancer--properties--client_side_defense--policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [client_side_defense](data-sources--http_loadbalancer--properties--client_side_defense.md)
- client_side_defense.policy

<a id="section"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

## Direct properties

- [disable_js_insert](data-sources--http_loadbalancer--properties--client_side_defense--policy--disable_js_insert.md): complete subsection reference.

- [js_insert_all_pages](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.disable_js_insert](data-sources--http_loadbalancer--properties--client_side_defense--policy--disable_js_insert.md)
- [client_side_defense.policy.js_insert_all_pages](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages.md)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except.md)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [client_side_defense](data-sources--http_loadbalancer--properties--client_side_defense.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
