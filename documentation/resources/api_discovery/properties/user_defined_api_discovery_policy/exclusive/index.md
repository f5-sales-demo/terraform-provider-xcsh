---
page_title: "user_defined_api_discovery_policy.exclusive"
subcategory: ""
description: "Configuration for exclusion action."
xcsh_docs: {"aliases": ["user defined api discovery policy exclusive"], "body_bytes": 2372, "body_sha256": "sha256:e76bed02d12a789178567d6e2bb573531e662d335adf73f4cb8564601f0e87cc", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "path": "documentation/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1323111022121002-1132133223122101-2223013020320022-0231220000301312-3210230203331032-2003110131131101-2120100303301100-3201211002333002", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.exclusive:ConflictingObjectAttributes:archive,ignore", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.exclusive:ConflictingObjectAttributes:archive,ignore", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "schema_version": 1, "sections": [{"aliases": ["user defined api discovery policy exclusive archive"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive", "archive"], "syntax": "attribute", "type": "object"}, {"aliases": ["user defined api discovery policy exclusive ignore"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive", "ignore"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration for exclusion action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.exclusive

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/)
- user_defined_api_discovery_policy.exclusive

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("archive",
    "ignore")}
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
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

Terraform syntax:

```terraform
exclusive {
  # Configure direct properties listed below.
}
```

## Direct properties

- [archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/archive/): complete subsection reference.

- [ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.exclusive.archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/archive/)
- [user_defined_api_discovery_policy.exclusive.ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
