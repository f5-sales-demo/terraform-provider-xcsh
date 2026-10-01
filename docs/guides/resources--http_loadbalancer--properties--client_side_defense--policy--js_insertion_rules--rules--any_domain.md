---
page_title: "client_side_defense.policy.js_insertion_rules.rules.any_domain"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules.rules.any_domain for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1585, "body_sha256": "sha256:9a8ff983bdd39d0dceee443527996afb8f3fc6c311257e4b1490ad3ca1df6e5a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "path": "docs/guides/resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--any_domain.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "any_domain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/any_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules.rules.any_domain for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.rules.any_domain

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [client_side_defense](resources--http_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](resources--http_loadbalancer--properties--client_side_defense--policy.md)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

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

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
