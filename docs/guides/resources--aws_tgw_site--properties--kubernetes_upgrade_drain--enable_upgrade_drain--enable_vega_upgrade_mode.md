---
page_title: "kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode"
subcategory: ""
description: "kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1343, "body_sha256": "sha256:e5e0a8ff847a06e95a0b90eb5ae976199716101656b2b475da30e02ded0aae43", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "path": "docs/guides/resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "enable_vega_upgrade_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain.md)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
