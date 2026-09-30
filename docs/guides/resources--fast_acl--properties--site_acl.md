---
page_title: "site_acl"
subcategory: ""
description: "site_acl for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2557, "body_sha256": "sha256:94f5cff28d39d8bddeb474c9648625bbc37358e85c39c07f6a304d03d7bcb9b6", "canonical_id": "xcsh-docs:resources:fast_acl:properties:site_acl", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:all_services", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "xcsh-docs:resources:fast_acl:properties:site_acl:inside_network", "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "xcsh-docs:resources:fast_acl:properties:site_acl:outside_network", "xcsh-docs:resources:fast_acl:properties:site_acl:vip_services"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl", "parent_id": "xcsh-docs:resources:fast_acl:reference", "path": "docs/guides/resources--fast_acl--properties--site_acl.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_acl

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- site_acl

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_services",
    "interface_services"),
  validators.ConflictingObjectAttributes("all_services",
    "vip_services"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("interface_services",
    "vip_services")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

Terraform syntax:

```terraform
site_acl {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_services](resources--fast_acl--properties--site_acl--all_services.md): complete subsection reference.

- [fast_acl_rules](resources--fast_acl--properties--site_acl--fast_acl_rules.md): complete subsection reference.

- [inside_network](resources--fast_acl--properties--site_acl--inside_network.md): complete subsection reference.

- [interface_services](resources--fast_acl--properties--site_acl--interface_services.md): complete subsection reference.

- [outside_network](resources--fast_acl--properties--site_acl--outside_network.md): complete subsection reference.

- [vip_services](resources--fast_acl--properties--site_acl--vip_services.md): complete subsection reference.

## Next pages

- [site_acl.all_services](resources--fast_acl--properties--site_acl--all_services.md)
- [site_acl.fast_acl_rules](resources--fast_acl--properties--site_acl--fast_acl_rules.md)
- [site_acl.inside_network](resources--fast_acl--properties--site_acl--inside_network.md)
- [site_acl.interface_services](resources--fast_acl--properties--site_acl--interface_services.md)
- [site_acl.outside_network](resources--fast_acl--properties--site_acl--outside_network.md)
- [site_acl.vip_services](resources--fast_acl--properties--site_acl--vip_services.md)
- [Property reference](resources--fast_acl--reference.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
