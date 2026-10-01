---
page_title: "disable_https_management"
subcategory: ""
description: "disable_https_management for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1324, "body_sha256": "sha256:cd5b12eea3b1c7383059954ee6c539c17ec93f359710117e725b7013bbfc3fa8", "canonical_id": "xcsh-docs:resources:nfv_service:properties:disable_https_management", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:disable_https_management", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "docs/guides/resources--nfv_service--properties--disable_https_management.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_https_management"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/disable_https_management/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_https_management for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_https_management

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- disable_https_management

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_https\_management, https\_management; Default: disable\_https\_management\]
Configuration parameter for disable https management.

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

OneOf alternatives in this subsection:

- [disable_https_management](resources--nfv_service--properties--disable_https_management.md#section)
- [https_management](resources--nfv_service--properties--https_management.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_https_management = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--nfv_service--reference.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
