---
page_title: "re_acl.selected_tenant_vip"
subcategory: ""
description: "Select various tenant public VIP(s)"
xcsh_docs: {"aliases": ["re acl selected tenant vip"], "body_bytes": 2128, "body_sha256": "sha256:9daf76c0036d3879e8200279c80c08a1d95d89b4a44c0a75cea2f22fb768e305", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl", "path": "documentation/resources/fast_acl/properties/re_acl/selected_tenant_vip/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2313112310232031-0132231301303321-1222013103002123-1203001002202003-2121231321230200-2210112100123321-0011302310000002-2202122102102002", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.selected_tenant_vip:RequiredObjectAttributes:public_ip_refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "selected_tenant_vip"], "schema_version": 1, "sections": [{"aliases": ["default tenant vip"], "anchor": "schema-re_acl--selected_tenant_vip--default_tenant_vip", "description": "Include tenant VIP in list of specific VIP(s)", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "selected_tenant_vip", "default_tenant_vip"], "syntax": "attribute", "type": "bool"}, {"aliases": ["public ip refs"], "anchor": "section", "description": "Select additional public VIP(s)", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-re_acl--selected_tenant_vip--public_ip_refs--name", "enforcement": "provider-schema", "group": "re_acl.selected_tenant_vip.public_ip_refs:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs", "type": "requires"}], "schema_path": ["re_acl", "selected_tenant_vip", "public_ip_refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/selected_tenant_vip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select various tenant public VIP(s)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.selected_tenant_vip

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- re_acl.selected_tenant_vip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specific Tenant VIP. Select various tenant public VIP(s)

Upstream description:

Select various tenant public VIP(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("public_ip_refs")}
```

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
selected_tenant_vip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-re_acl--selected_tenant_vip--default_tenant_vip"></a>

### default_tenant_vip property

Type: `"bool"`. Optional.

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

- [public_ip_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/selected_tenant_vip/public_ip_refs/): complete subsection reference.

## Next pages

- [re_acl.selected_tenant_vip.public_ip_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/selected_tenant_vip/public_ip_refs/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
