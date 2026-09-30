---
page_title: "re_acl"
subcategory: ""
description: "re_acl for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2280, "body_sha256": "sha256:916f5973a55978a77a3ee626d4a44f5da8e1e38dc2cdab64927942406cbd10b7", "canonical_id": "xcsh-docs:resources:fast_acl:properties:re_acl", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl", "parent_id": "xcsh-docs:resources:fast_acl:reference", "path": "docs/guides/resources--fast_acl--properties--re_acl.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# re_acl

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- re_acl

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_public_vips",
    "default_tenant_vip"),
  validators.ConflictingObjectAttributes("all_public_vips",
    "selected_tenant_vip"),
  validators.ConflictingObjectAttributes("default_tenant_vip",
    "selected_tenant_vip")}
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
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](resources--fast_acl--properties--re_acl.md#section)
- [site_acl](resources--fast_acl--properties--site_acl.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
re_acl {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_public_vips](resources--fast_acl--properties--re_acl--all_public_vips.md): complete subsection reference.

- [default_tenant_vip](resources--fast_acl--properties--re_acl--default_tenant_vip.md): complete subsection reference.

- [fast_acl_rules](resources--fast_acl--properties--re_acl--fast_acl_rules.md): complete subsection reference.

- [selected_tenant_vip](resources--fast_acl--properties--re_acl--selected_tenant_vip.md): complete subsection reference.

## Next pages

- [re_acl.all_public_vips](resources--fast_acl--properties--re_acl--all_public_vips.md)
- [re_acl.default_tenant_vip](resources--fast_acl--properties--re_acl--default_tenant_vip.md)
- [re_acl.fast_acl_rules](resources--fast_acl--properties--re_acl--fast_acl_rules.md)
- [re_acl.selected_tenant_vip](resources--fast_acl--properties--re_acl--selected_tenant_vip.md)
- [Property reference](resources--fast_acl--reference.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
