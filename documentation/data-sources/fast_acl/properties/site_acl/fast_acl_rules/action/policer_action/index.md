---
page_title: "site_acl.fast_acl_rules.action.policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["site acl fast acl rules action policer action"], "body_bytes": 1314, "body_sha256": "sha256:b87186ee745ce88c47cd97a80d8efb48023070d3f9fd2bde3401951ddf6aeecc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "path": "documentation/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1000110003033231-0123100213212023-2213112112021320-1001123013012000-3323311311022333-2101201032033201-3021301213302113-1001220323323332", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "action", "policer_action"], "schema_version": 1, "sections": [{"aliases": ["site acl fast acl rules action policer action ref"], "anchor": "section", "description": "A policer direct reference.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action", "policer_action", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_aclCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.action.policer_action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/)
- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- site_acl.fast_acl_rules.action.policer_action

<a id="section"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

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

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/policer_action/ref/): complete subsection reference.
