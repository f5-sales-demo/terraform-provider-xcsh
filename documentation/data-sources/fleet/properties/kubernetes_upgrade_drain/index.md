---
page_title: "kubernetes_upgrade_drain"
subcategory: ""
description: "kubernetes_upgrade_drain for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1770, "body_sha256": "sha256:d6a67fc85838ea96f23f22c5cd83d3fedf0dc137ccf0179bfcb72cf799d95ec7", "child_ids": ["xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/kubernetes_upgrade_drain/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["kubernetes_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
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

- [disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/kubernetes_upgrade_drain/disable_upgrade_drain/): complete subsection reference.

- [enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/): complete subsection reference.

## Next pages

- [kubernetes_upgrade_drain.disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/kubernetes_upgrade_drain/disable_upgrade_drain/)
- [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
