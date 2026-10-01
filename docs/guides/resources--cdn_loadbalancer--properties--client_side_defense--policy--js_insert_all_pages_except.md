---
page_title: "client_side_defense.policy.js_insert_all_pages_except"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insert_all_pages_except for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1532, "body_sha256": "sha256:c9f8bb3b33e793b5821780d09d658bc4ff6d8f0142949b0db2d029e0a996600d", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy", "path": "docs/guides/resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insert_all_pages_except for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insert_all_pages_except

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [client_side_defense](resources--cdn_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](resources--cdn_loadbalancer--properties--client_side_defense--policy.md)
- client_side_defense.policy.js_insert_all_pages_except

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_list](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except--exclude_list.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except--exclude_list.md)
- [client_side_defense.policy](resources--cdn_loadbalancer--properties--client_side_defense--policy.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
