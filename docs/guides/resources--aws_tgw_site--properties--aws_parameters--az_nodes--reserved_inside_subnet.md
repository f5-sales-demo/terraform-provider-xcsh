---
page_title: "aws_parameters.az_nodes.reserved_inside_subnet"
subcategory: ""
description: "aws_parameters.az_nodes.reserved_inside_subnet for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1183, "body_sha256": "sha256:26bbbea615bc199d1bab2c28a4c42c7e76fd9dc3322ad9e6660675263854cac2", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "path": "docs/guides/resources--aws_tgw_site--properties--aws_parameters--az_nodes--reserved_inside_subnet.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_parameters", "az_nodes", "reserved_inside_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/az_nodes/reserved_inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.az_nodes.reserved_inside_subnet for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.az_nodes.reserved_inside_subnet

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
- [aws_parameters.az_nodes](resources--aws_tgw_site--properties--aws_parameters--az_nodes.md)
- aws_parameters.az_nodes.reserved_inside_subnet

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved inside subnet.

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
reserved_inside_subnet = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_parameters.az_nodes](resources--aws_tgw_site--properties--aws_parameters--az_nodes.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
