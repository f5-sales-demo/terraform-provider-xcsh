---
page_title: "user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion"
subcategory: ""
description: "Configuration for exclusion action."
xcsh_docs: {"aliases": ["user defined api discovery policy discovery rules rule properties exclusion"], "body_bytes": 2870, "body_sha256": "sha256:ac9f8cd0cfda241a7da6958caa7f6d9fbeeec83f67e1e7a6d9c61f9b0a95cff4", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion:archive", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion:ignore"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "exclusion"], "schema_version": 1, "sections": [{"aliases": ["archive"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion:archive", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "exclusion", "archive"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion:ignore", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "exclusion", "ignore"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for exclusion action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/)
- [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

<a id="section"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

## Direct properties

- [archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/archive/): complete subsection reference.

- [ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/ignore/): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/archive/)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/ignore/)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
