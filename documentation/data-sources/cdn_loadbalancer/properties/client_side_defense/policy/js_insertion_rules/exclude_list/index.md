---
page_title: "client_side_defense.policy.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules exclude list"], "body_bytes": 3987, "body_sha256": "sha256:db5f117551aed2f9f80ccc9070c950b44350981f8c0348a6c4de3f2df4cfe267", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "path": "documentation/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insertion rules exclude list any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules exclude list domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules exclude list metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules exclude list path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list", "path"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/path/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/any_domain/)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/domain/)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/metadata/)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/path/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
