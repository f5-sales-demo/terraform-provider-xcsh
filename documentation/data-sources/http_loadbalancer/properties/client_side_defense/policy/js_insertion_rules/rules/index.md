---
page_title: "client_side_defense.policy.js_insertion_rules.rules"
subcategory: "Load Balancing"
description: "Required list of pages to insert Client-Side Defense client JavaScript."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules rules"], "body_bytes": 2921, "body_sha256": "sha256:4d4013ade55e9c9bc920b9a5e8a5aa6a80ee1a1654b73f855962aa739eebcb04", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "documentation/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insertion rules rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules rules domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules rules path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "path"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Required list of pages to insert Client-Side Defense client JavaScript.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- client_side_defense.policy.js_insertion_rules.rules

<a id="section"></a>

Type: `"list"`. Computed.

Required list of pages to insert Client-Side Defense client JavaScript.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/path/): complete subsection reference.
