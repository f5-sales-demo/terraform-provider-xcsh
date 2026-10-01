---
page_title: "site_acl"
subcategory: ""
description: "site_acl for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2970, "body_sha256": "sha256:a61787331135466f9ac4d0876dd79b2224e06625ca42bb36871071488378362a", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:site_acl:all_services", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "xcsh-docs:data-sources:fast_acl:properties:site_acl:inside_network", "xcsh-docs:data-sources:fast_acl:properties:site_acl:interface_services", "xcsh-docs:data-sources:fast_acl:properties:site_acl:outside_network", "xcsh-docs:data-sources:fast_acl:properties:site_acl:vip_services"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl", "parent_id": "xcsh-docs:data-sources:fast_acl:reference", "path": "documentation/data-sources/fast_acl/properties/site_acl/index.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["site_acl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- site_acl

<a id="section"></a>

Type: `"single"`. Computed.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

## Direct properties

- [all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/all_services/): complete subsection reference.

- [fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/): complete subsection reference.

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/inside_network/): complete subsection reference.

- [interface_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/interface_services/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/outside_network/): complete subsection reference.

- [vip_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/vip_services/): complete subsection reference.

## Next pages

- [site_acl.all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/all_services/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/)
- [site_acl.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/inside_network/)
- [site_acl.interface_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/interface_services/)
- [site_acl.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/outside_network/)
- [site_acl.vip_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/vip_services/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
