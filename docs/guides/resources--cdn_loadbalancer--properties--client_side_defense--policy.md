---
page_title: "client_side_defense.policy"
subcategory: "Load Balancing"
description: "client_side_defense.policy for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3092, "body_sha256": "sha256:cda5925905cdf54057fe5708bf60134b023f73eb6eb99e7f502a80a315b63184", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense", "path": "docs/guides/resources--cdn_loadbalancer--properties--client_side_defense--policy.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [client_side_defense](resources--cdn_loadbalancer--properties--client_side_defense.md)
- client_side_defense.policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
```

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

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_js_insert](resources--cdn_loadbalancer--properties--client_side_defense--policy--disable_js_insert.md): complete subsection reference.

- [js_insert_all_pages](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.disable_js_insert](resources--cdn_loadbalancer--properties--client_side_defense--policy--disable_js_insert.md)
- [client_side_defense.policy.js_insert_all_pages](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages.md)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insert_all_pages_except.md)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [client_side_defense](resources--cdn_loadbalancer--properties--client_side_defense.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
