---
page_title: "direct_connect_enabled.standard_vifs"
subcategory: ""
description: "direct_connect_enabled.standard_vifs for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1064, "body_sha256": "sha256:870e7966e42e510b947fc315fb29772ece99a3b77be2d0286adf572ed9908317", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:standard_vifs", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:standard_vifs", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "path": "docs/guides/resources--aws_tgw_site--properties--direct_connect_enabled--standard_vifs.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["direct_connect_enabled", "standard_vifs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/standard_vifs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled.standard_vifs for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.standard_vifs

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md)
- direct_connect_enabled.standard_vifs

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for standard vifs.

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
standard_vifs = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [direct_connect_enabled](resources--aws_tgw_site--properties--direct_connect_enabled.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
