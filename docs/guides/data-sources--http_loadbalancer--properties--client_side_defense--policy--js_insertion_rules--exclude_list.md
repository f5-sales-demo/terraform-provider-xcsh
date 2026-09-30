---
page_title: "client_side_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules.exclude_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3175, "body_sha256": "sha256:35ef7932174be4e205aed0fc58d4b8683f9a1cc8c4c35cfa8b4203ccda6c4f69", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules.exclude_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# client_side_defense.policy.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [client_side_defense](data-sources--http_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](data-sources--http_loadbalancer--properties--client_side_defense--policy.md)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- client_side_defense.policy.js_insertion_rules.exclude_list

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

- [any_domain](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--any_domain.md): complete subsection reference.

- [domain](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--domain.md): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--metadata.md): complete subsection reference.

- [path](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--path.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--any_domain.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--domain.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--metadata.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--path.md)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
