---
page_title: "re_acl.selected_tenant_vip"
subcategory: ""
description: "re_acl.selected_tenant_vip for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1513, "body_sha256": "sha256:4ccbc9ed8af2799a417e78ba6ee9a83b9060f46d2266a9f097b411b7e41c54c7", "canonical_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "path": "docs/guides/data-sources--fast_acl--properties--re_acl--selected_tenant_vip.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl", "selected_tenant_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl.selected_tenant_vip for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.selected_tenant_vip

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md)
- [Property reference](data-sources--fast_acl--reference.md)
- [re_acl](data-sources--fast_acl--properties--re_acl.md)
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

- [public_ip_refs](data-sources--fast_acl--properties--re_acl--selected_tenant_vip--public_ip_refs.md): complete subsection reference.

## Next pages

- [re_acl.selected_tenant_vip.public_ip_refs](data-sources--fast_acl--properties--re_acl--selected_tenant_vip--public_ip_refs.md)
- [re_acl](data-sources--fast_acl--properties--re_acl.md)
- [xcsh_fast_acl](../data-sources/fast_acl.md)
