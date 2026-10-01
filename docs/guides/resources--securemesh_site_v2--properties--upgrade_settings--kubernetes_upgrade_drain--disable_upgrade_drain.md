---
page_title: "upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain"
subcategory: ""
description: "upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1339, "body_sha256": "sha256:c0f811a9c3b47966ca052dfd070869cec97a4b61707ea2a4d01f70c731d32d8b", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:disable_upgrade_drain", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:disable_upgrade_drain", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "path": "docs/guides/resources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain--disable_upgrade_drain.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "disable_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/disable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [upgrade_settings](resources--securemesh_site_v2--properties--upgrade_settings.md)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain.md)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

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

- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
