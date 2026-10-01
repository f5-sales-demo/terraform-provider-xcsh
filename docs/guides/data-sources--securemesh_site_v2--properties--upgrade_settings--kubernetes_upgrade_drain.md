---
page_title: "upgrade_settings.kubernetes_upgrade_drain"
subcategory: ""
description: "upgrade_settings.kubernetes_upgrade_drain for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1737, "body_sha256": "sha256:0cb2f8f362a79b8f1d4d459eff668bc3ba4f6bfcfd9fb8cf96e936ae77a80f8f", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings", "path": "docs/guides/data-sources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upgrade_settings.kubernetes_upgrade_drain for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upgrade_settings.kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [upgrade_settings](data-sources--securemesh_site_v2--properties--upgrade_settings.md)
- upgrade_settings.kubernetes_upgrade_drain

<a id="section"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

## Direct properties

- [disable_upgrade_drain](data-sources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain--disable_upgrade_drain.md): complete subsection reference.

- [enable_upgrade_drain](data-sources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain.md): complete subsection reference.

## Next pages

- [upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain--disable_upgrade_drain.md)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--properties--upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- [upgrade_settings](data-sources--securemesh_site_v2--properties--upgrade_settings.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
