---
page_title: "kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode"
subcategory: "Infrastructure"
description: "kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1248, "body_sha256": "sha256:35e89d25a1bfe7303cc19df65aa2da96d7cf8d404e637971443e39ad41d5a05c", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "path": "docs/guides/resources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "disable_vega_upgrade_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--properties--kubernetes_upgrade_drain.md)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
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

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
