---
page_title: "client_side_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules.exclude_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3518, "body_sha256": "sha256:2108c30dfd64f938d67409dee16d48fb84450a8ca4611b39d8c3291be2ae7829", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "docs/guides/resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules.exclude_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [client_side_defense](resources--http_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](resources--http_loadbalancer--properties--client_side_defense--policy.md)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--any_domain.md): complete subsection reference.

- [domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--domain.md): complete subsection reference.

- [metadata](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--metadata.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--path.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--any_domain.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--domain.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--metadata.md)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--exclude_list--path.md)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
