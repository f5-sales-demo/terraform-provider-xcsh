---
page_title: "re_acl"
subcategory: ""
description: "re_acl for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2654, "body_sha256": "sha256:04d2697e10ea48ae3ed4510280a3cc56cf5049171097fa26f942dced98a3f0df", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:all_public_vips", "xcsh-docs:data-sources:fast_acl:properties:re_acl:default_tenant_vip", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "xcsh-docs:data-sources:fast_acl:properties:re_acl:selected_tenant_vip"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl", "parent_id": "xcsh-docs:data-sources:fast_acl:reference", "path": "documentation/data-sources/fast_acl/properties/re_acl/index.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["re_acl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/#section)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [all_public_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/all_public_vips/): complete subsection reference.

- [default_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/default_tenant_vip/): complete subsection reference.

- [fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/): complete subsection reference.

- [selected_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/): complete subsection reference.

## Next pages

- [re_acl.all_public_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/all_public_vips/)
- [re_acl.default_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/default_tenant_vip/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/)
- [re_acl.selected_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/selected_tenant_vip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
