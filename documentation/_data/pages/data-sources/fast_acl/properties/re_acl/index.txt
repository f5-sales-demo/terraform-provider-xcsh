---
page_title: "re_acl"
subcategory: ""
description: "Fast ACL definition for RE."
xcsh_docs: {"aliases": ["re acl"], "body_bytes": 1798, "body_sha256": "sha256:f202e39cd8389e19abd41fa2b5acf624a030e1bf6cdc17236f7ecb6b25250eb4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:all_public_vips", "xcsh-docs:data-sources:fast_acl:properties:re_acl:default_tenant_vip", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "parent_id": "xcsh-docs:data-sources:fast_acl:reference", "path": "documentation/data-sources/fast_acl/properties/re_acl/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl"], "schema_version": 1, "sections": [{"aliases": ["re acl all public vips"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:all_public_vips", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "all_public_vips"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl default tenant vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:default_tenant_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "default_tenant_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules"], "anchor": "section", "description": "Fast ACL rules to match.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl selected tenant vip"], "anchor": "section", "description": "Select various tenant public VIP(s)", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "selected_tenant_vip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Fast ACL definition for RE.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["fast_aclCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- re_acl

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/#section)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [all_public_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/all_public_vips/): complete subsection reference.

- [default_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/default_tenant_vip/): complete subsection reference.

- [fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/): complete subsection reference.

- [selected_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/): complete subsection reference.
