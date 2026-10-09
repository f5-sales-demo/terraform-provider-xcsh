---
page_title: "client_side_defense.policy"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Client-Side Defense policy."
xcsh_docs: {"aliases": ["client side defense policy"], "body_bytes": 1944, "body_sha256": "sha256:c4a4ca29fa1290d87ac8773b59f677f6eb84928b96aa5639c1507f03587d6621", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense", "path": "documentation/resources/http_loadbalancer/properties/client_side_defense/policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages except"], "anchor": "section", "description": "Insert Client-Side Defense JavaScript in all pages with the exceptions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except"], "syntax": "block", "type": "object"}, {"aliases": ["client side defense policy js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Client-Side Defense Policy.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This defines various configuration OPTIONS for Client-Side Defense policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
