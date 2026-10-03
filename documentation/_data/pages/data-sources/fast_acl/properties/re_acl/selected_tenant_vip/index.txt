---
page_title: "re_acl.selected_tenant_vip"
subcategory: ""
description: "Select various tenant public VIP(s)"
xcsh_docs: {"aliases": ["re acl selected tenant vip"], "body_bytes": 1868, "body_sha256": "sha256:4e97add0801fd26663d04fc5931095ec77e98520fa0f176301e3cd7926dda5eb", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "path": "documentation/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1210100311322300-1002001032202203-0012331030302121-3022301020332021-1322201112200222-3300122330231313-3030120323233302-3001322221233201", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "selected_tenant_vip"], "schema_version": 1, "sections": [{"aliases": ["re acl selected tenant vip default tenant vip"], "anchor": "schema-re_acl--selected_tenant_vip--default_tenant_vip", "description": "Include tenant VIP in list of specific VIP(s)", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "selected_tenant_vip", "default_tenant_vip"], "syntax": "attribute", "type": "bool"}, {"aliases": ["re acl selected tenant vip public ip refs"], "anchor": "section", "description": "Select additional public VIP(s)", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["re_acl", "selected_tenant_vip", "public_ip_refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Select various tenant public VIP(s)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.selected_tenant_vip

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/)
- re_acl.selected_tenant_vip

<a id="section"></a>

Type: `"single"`. Computed.

Specific Tenant VIP. Select various tenant public VIP(s)

Upstream description:

Select various tenant public VIP(s)

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

<a id="schema-re_acl--selected_tenant_vip--default_tenant_vip"></a>

### default_tenant_vip property

Type: `"bool"`. Computed.

Include tenant VIP in list of specific VIP(s).

Upstream description:

Include tenant VIP in list of specific VIP(s)

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

- [public_ip_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/public_ip_refs/): complete subsection reference.

## Next pages

- [re_acl.selected_tenant_vip.public_ip_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/public_ip_refs/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
