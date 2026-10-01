---
page_title: "custom_storage_config.no_storage_interfaces"
subcategory: ""
description: "custom_storage_config.no_storage_interfaces for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1104, "body_sha256": "sha256:982cb076f5338ff63236d36582909e28a26d4b0e69d9d7490b2e1c20a1a87c7b", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_storage_interfaces", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_storage_interfaces", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--no_storage_interfaces.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "no_storage_interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/no_storage_interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.no_storage_interfaces for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.no_storage_interfaces

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- custom_storage_config.no_storage_interfaces

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no storage interfaces.

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
no_storage_interfaces = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
