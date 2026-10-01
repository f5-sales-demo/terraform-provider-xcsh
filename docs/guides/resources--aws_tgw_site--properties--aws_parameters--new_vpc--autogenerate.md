---
page_title: "aws_parameters.new_vpc.autogenerate"
subcategory: ""
description: "aws_parameters.new_vpc.autogenerate for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1137, "body_sha256": "sha256:d9c83f0c9b5059d0a69bdb7a726734c839f06d64d512ae99ecdb45dde5ad765c", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "path": "docs/guides/resources--aws_tgw_site--properties--aws_parameters--new_vpc--autogenerate.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_parameters", "new_vpc", "autogenerate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_vpc/autogenerate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.new_vpc.autogenerate for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_vpc.autogenerate

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
- [aws_parameters.new_vpc](resources--aws_tgw_site--properties--aws_parameters--new_vpc.md)
- aws_parameters.new_vpc.autogenerate

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for autogenerate.

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
autogenerate = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_parameters.new_vpc](resources--aws_tgw_site--properties--aws_parameters--new_vpc.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
