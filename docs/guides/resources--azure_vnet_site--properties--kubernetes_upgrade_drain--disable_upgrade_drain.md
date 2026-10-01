---
page_title: "kubernetes_upgrade_drain.disable_upgrade_drain"
subcategory: "Infrastructure"
description: "kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1129, "body_sha256": "sha256:be8a0b30d76fcb5d3f56a4959ece35e89988dd04b1802355de7b192985c405e8", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:kubernetes_upgrade_drain", "path": "docs/guides/resources--azure_vnet_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.disable_upgrade_drain

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [kubernetes_upgrade_drain](resources--azure_vnet_site--properties--kubernetes_upgrade_drain.md)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kubernetes_upgrade_drain](resources--azure_vnet_site--properties--kubernetes_upgrade_drain.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
