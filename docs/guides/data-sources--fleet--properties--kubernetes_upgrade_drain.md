---
page_title: "kubernetes_upgrade_drain"
subcategory: ""
description: "kubernetes_upgrade_drain for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1362, "body_sha256": "sha256:ec9669fc61347557c80ee46e6aae740fff39a51db8eb35ca1174a04ba760da7e", "canonical_id": "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain", "child_ids": ["xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "docs/guides/data-sources--fleet--properties--kubernetes_upgrade_drain.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- kubernetes_upgrade_drain

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

- [disable_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md): complete subsection reference.

- [enable_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md): complete subsection reference.

## Next pages

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- [Property reference](data-sources--fleet--reference.md)
- [xcsh_fleet](../data-sources/fleet.md)
