---
page_title: "client_side_defense.policy.js_insertion_rules.rules"
subcategory: "Load Balancing"
description: "client_side_defense.policy.js_insertion_rules.rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3553, "body_sha256": "sha256:d8621a316e3b8bcc84e9d6484cf64701f3f467e8ec403a3287e17cbd6d879f48", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "docs/guides/resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense.policy.js_insertion_rules.rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# client_side_defense.policy.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [client_side_defense](resources--http_loadbalancer--properties--client_side_defense.md)
- [client_side_defense.policy](resources--http_loadbalancer--properties--client_side_defense--policy.md)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- client_side_defense.policy.js_insertion_rules.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--any_domain.md): complete subsection reference.

- [domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--domain.md): complete subsection reference.

- [metadata](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--metadata.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--path.md): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--any_domain.md)
- [client_side_defense.policy.js_insertion_rules.rules.domain](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--domain.md)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--metadata.md)
- [client_side_defense.policy.js_insertion_rules.rules.path](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules--rules--path.md)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--client_side_defense--policy--js_insertion_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
