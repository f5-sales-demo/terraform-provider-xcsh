---
page_title: "user_defined_api_discovery_policy"
subcategory: ""
description: "Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field."
xcsh_docs: {"aliases": ["user defined api discovery policy"], "body_bytes": 2601, "body_sha256": "sha256:ad331ffc7b221c1f559d9269c8570fb7fe9ba3f1b0b3c9881d9de93fb1a9a3e4", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "parent_id": "xcsh-docs:resources:api_discovery:reference", "path": "documentation/resources/api_discovery/properties/user_defined_api_discovery_policy/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2110113123313300-1200313003201203-0030202211112133-3211202003002122-0322210303301001-2132132313312231-2300101120203001-1102330022023122", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy:ConflictingObjectAttributes:exclusive,inclusive", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy:ConflictingObjectAttributes:exclusive,inclusive", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy"], "schema_version": 1, "sections": [{"aliases": ["user defined api discovery policy discovery rules"], "anchor": "section", "description": "Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom; unmatched endpoints follow the default action.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules"], "syntax": "block", "type": "object"}, {"aliases": ["user defined api discovery policy exclusive"], "anchor": "section", "description": "Configuration for exclusion action.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.exclusive:ConflictingObjectAttributes:archive,ignore", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.exclusive:ConflictingObjectAttributes:archive,ignore", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore", "type": "conflicts"}], "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "syntax": "block", "type": "object"}, {"aliases": ["user defined api discovery policy inclusive"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "inclusive"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- user_defined_api_discovery_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exclusive",
    "inclusive")}
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
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

Terraform syntax:

```terraform
user_defined_api_discovery_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/): complete subsection reference.

- [exclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/): complete subsection reference.

- [inclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/)
- [user_defined_api_discovery_policy.exclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/)
- [user_defined_api_discovery_policy.inclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
