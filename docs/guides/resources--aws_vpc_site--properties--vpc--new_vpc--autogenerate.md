---
page_title: "vpc.new_vpc.autogenerate"
subcategory: "Infrastructure"
description: "vpc.new_vpc.autogenerate for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 950, "body_sha256": "sha256:566aebc424fe32f068d2ddf298bd2cff96987d2f29aa8409def8d6e7d180c2ec", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "path": "docs/guides/resources--aws_vpc_site--properties--vpc--new_vpc--autogenerate.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vpc", "new_vpc", "autogenerate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/vpc/new_vpc/autogenerate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc.new_vpc.autogenerate for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vpc.new_vpc.autogenerate

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [vpc](resources--aws_vpc_site--properties--vpc.md)
- [vpc.new_vpc](resources--aws_vpc_site--properties--vpc--new_vpc.md)
- vpc.new_vpc.autogenerate

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

- [vpc.new_vpc](resources--aws_vpc_site--properties--vpc--new_vpc.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
