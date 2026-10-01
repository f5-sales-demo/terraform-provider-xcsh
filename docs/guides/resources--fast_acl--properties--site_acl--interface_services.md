---
page_title: "site_acl.interface_services"
subcategory: ""
description: "site_acl.interface_services for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 972, "body_sha256": "sha256:ef449d47a383baa4ddf4e8f426e33ab62653d6638317b54254d86b138cb68819", "canonical_id": "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "child_ids": [], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl", "path": "docs/guides/resources--fast_acl--properties--site_acl--interface_services.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl", "interface_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/interface_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl.interface_services for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.interface_services

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [site_acl](resources--fast_acl--properties--site_acl.md)
- site_acl.interface_services

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for interface services.

Upstream description:

This can be used for messages where no values are needed.

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
interface_services = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [site_acl](resources--fast_acl--properties--site_acl.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
