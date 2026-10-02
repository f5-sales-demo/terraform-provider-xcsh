---
page_title: "user_defined_api_discovery_policy"
subcategory: ""
description: "Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field."
xcsh_docs: {"aliases": ["user defined api discovery policy"], "body_bytes": 2321, "body_sha256": "sha256:14b1d13bb46be214972bfb77e7ce7f4867add71b9215c98a6775ce1aab094e31", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:inclusive"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy", "parent_id": "xcsh-docs:data-sources:api_discovery:reference", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy"], "schema_version": 1, "sections": [{"aliases": ["discovery rules"], "anchor": "section", "description": "Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom; unmatched endpoints follow the default action.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["exclusive"], "anchor": "section", "description": "Configuration for exclusion action.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "syntax": "attribute", "type": "object"}, {"aliases": ["inclusive"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "inclusive"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- user_defined_api_discovery_policy

<a id="section"></a>

Type: `"single"`. Computed.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

## Direct properties

- [discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/): complete subsection reference.

- [exclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/): complete subsection reference.

- [inclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/)
- [user_defined_api_discovery_policy.exclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/)
- [user_defined_api_discovery_policy.inclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
