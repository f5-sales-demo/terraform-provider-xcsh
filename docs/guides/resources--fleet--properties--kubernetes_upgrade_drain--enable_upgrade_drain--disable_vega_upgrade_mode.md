---
page_title: "kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode"
subcategory: ""
description: "kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1291, "body_sha256": "sha256:c48c9b8e8c4c5d8c2d830c6b144524ad5e94d0e7fb4077e1d2879878220ba42d", "canonical_id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "parent_id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "path": "docs/guides/resources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "disable_vega_upgrade_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [kubernetes_upgrade_drain](resources--fleet--properties--kubernetes_upgrade_drain.md)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- [xcsh_fleet](../resources/fleet.md)
