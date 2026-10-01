---
page_title: "ingress_egress_gw.no_dc_cluster_group"
subcategory: "Infrastructure"
description: "ingress_egress_gw.no_dc_cluster_group for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1028, "body_sha256": "sha256:20c88642dfbb0489bbe5433fc3f183084d416a5bbfa7204f826badd66caeb92d", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--no_dc_cluster_group.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "no_dc_cluster_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/no_dc_cluster_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.no_dc_cluster_group for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.no_dc_cluster_group

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.no_dc_cluster_group

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
no_dc_cluster_group = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
