---
page_title: "mum_action"
subcategory: ""
description: "Modify behavior for a matching request. The modification could be to entirely skip processing."
xcsh_docs: {"aliases": ["mum action"], "body_bytes": 1981, "body_sha256": "sha256:ffc4f16649324533f14c736093df7a740f50ba867d59533a226ebbc24b8e2fdb", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:mum_action:default", "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:mum_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/mum_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["mum_action"], "schema_version": 1, "sections": [{"aliases": ["mum action default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:default", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mum_action", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["mum action skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mum_action", "skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/mum_action/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Modify behavior for a matching request. The modification could be to entirely skip processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mum_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- mum_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default",
    "skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

Terraform syntax:

```terraform
mum_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/default/): complete subsection reference.

- [skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/skip_processing/): complete subsection reference.

## Next pages

- [mum_action.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/default/)
- [mum_action.skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/skip_processing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
