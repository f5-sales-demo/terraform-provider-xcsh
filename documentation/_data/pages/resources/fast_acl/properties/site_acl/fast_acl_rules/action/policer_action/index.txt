---
page_title: "site_acl.fast_acl_rules.action.policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["site acl fast acl rules action policer action"], "body_bytes": 1929, "body_sha256": "sha256:dc1e41bebfb6f0132c71ad915bc5de6da2a731b2f25d70e070385c39c48d1201", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action:ref"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "path": "documentation/resources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221203321120020-0122003030120223-1102012200130021-3323202212112111-1131302110020123-1330321231323101-2320010122110200-2201300021303001", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "action", "policer_action"], "schema_version": 1, "sections": [{"aliases": ["ref"], "anchor": "section", "description": "A policer direct reference.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action", "policer_action", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.action.policer_action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/)
- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- site_acl.fast_acl_rules.action.policer_action

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/ref/): complete subsection reference.

## Next pages

- [site_acl.fast_acl_rules.action.policer_action.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/ref/)
- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
