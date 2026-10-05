---
page_title: "re_acl.fast_acl_rules.action.policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["re acl fast acl rules action policer action"], "body_bytes": 1903, "body_sha256": "sha256:9fd527c63729b51a6d3db4980340d3c33a72f2d6e343e8fcf6f448383bd91785", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "path": "documentation/resources/fast_acl/properties/re_acl/fast_acl_rules/action/policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2323002023331031-1212132201020130-1133223202221102-2301310212213112-2030231023120022-0132213330112313-0001020121221131-2002310313033032", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "action", "policer_action"], "schema_version": 1, "sections": [{"aliases": ["re acl fast acl rules action policer action ref"], "anchor": "section", "description": "A policer direct reference.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "action", "policer_action", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/action/policer_action/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.action.policer_action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/)
- [re_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/)
- re_acl.fast_acl_rules.action.policer_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
policer_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/policer_action/ref/): complete subsection reference.

## Next pages

- [re_acl.fast_acl_rules.action.policer_action.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/policer_action/ref/)
- [re_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
