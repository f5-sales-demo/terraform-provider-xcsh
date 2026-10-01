---
page_title: "vpc_attachments"
subcategory: ""
description: "vpc_attachments for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1331, "body_sha256": "sha256:c7842583674039bde9ce4b7783506310bb60b8533268861b8f02bd819449d806", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/vpc_attachments/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["vpc_attachments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vpc_attachments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc_attachments for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc_attachments

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- vpc_attachments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Spoke VPCs to be attached to the AWS TGW Site.

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
vpc_attachments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/): complete subsection reference.

## Next pages

- [vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
