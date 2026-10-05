---
page_title: "client_side_defense.policy"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Client-Side Defense policy."
xcsh_docs: {"aliases": ["client side defense policy"], "body_bytes": 3756, "body_sha256": "sha256:4368945d8799fdfb3139fcab78852bd6188acb186a7f2c593b24bf54e81def3c", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense", "path": "documentation/resources/http_loadbalancer/properties/client_side_defense/policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages_except,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages_except,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages except"], "anchor": "section", "description": "Insert Client-Side Defense JavaScript in all pages with the exceptions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except"], "syntax": "block", "type": "object"}, {"aliases": ["client side defense policy js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Client-Side Defense Policy.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy.js_insertion_rules:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "type": "requires"}], "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines various configuration OPTIONS for Client-Side Defense policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/)
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

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/disable_js_insert/): complete subsection reference.

- [js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages/): complete subsection reference.

- [js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/): complete subsection reference.

## Next pages

- [client_side_defense.policy.disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/disable_js_insert/)
- [client_side_defense.policy.js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages/)
- [client_side_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
