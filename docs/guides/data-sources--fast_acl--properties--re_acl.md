---
page_title: "re_acl"
subcategory: ""
description: "re_acl for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1845, "body_sha256": "sha256:d4630410f406e7716d2fd9a8db65c547809cfa41d97c4ac9533985e9eea860ed", "canonical_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:all_public_vips", "xcsh-docs:data-sources:fast_acl:properties:re_acl:default_tenant_vip", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "parent_id": "xcsh-docs:data-sources:fast_acl:reference", "path": "docs/guides/data-sources--fast_acl--properties--re_acl.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# re_acl

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md)
- [Property reference](data-sources--fast_acl--reference.md)
- re_acl

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

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

- [re_acl](data-sources--fast_acl--properties--re_acl.md#section)
- [site_acl](data-sources--fast_acl--properties--site_acl.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [all_public_vips](data-sources--fast_acl--properties--re_acl--all_public_vips.md): complete subsection reference.

- [default_tenant_vip](data-sources--fast_acl--properties--re_acl--default_tenant_vip.md): complete subsection reference.

- [fast_acl_rules](data-sources--fast_acl--properties--re_acl--fast_acl_rules.md): complete subsection reference.

- [selected_tenant_vip](data-sources--fast_acl--properties--re_acl--selected_tenant_vip.md): complete subsection reference.

## Next pages

- [re_acl.all_public_vips](data-sources--fast_acl--properties--re_acl--all_public_vips.md)
- [re_acl.default_tenant_vip](data-sources--fast_acl--properties--re_acl--default_tenant_vip.md)
- [re_acl.fast_acl_rules](data-sources--fast_acl--properties--re_acl--fast_acl_rules.md)
- [re_acl.selected_tenant_vip](data-sources--fast_acl--properties--re_acl--selected_tenant_vip.md)
- [Property reference](data-sources--fast_acl--reference.md)
- [xcsh_fast_acl](../data-sources/fast_acl.md)
