---
page_title: "egress_gateway_default"
subcategory: "Infrastructure"
description: "egress_gateway_default for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1720, "body_sha256": "sha256:22996c2e66a0aaab01e53d1b66f82c0889663c2a7fabb9f2854f91a32fb5c59a", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:egress_gateway_default", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/egress_gateway_default/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["egress_gateway_default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/egress_gateway_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_gateway_default for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# egress_gateway_default

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- egress_gateway_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: egress\_gateway\_default, egress\_nat\_gw, egress\_virtual\_private\_gateway; Default:
egress\_gateway\_default\] Configuration parameter for egress gateway default.

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

OneOf alternatives in this subsection:

- [egress_gateway_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/egress_gateway_default/#section)
- [egress_nat_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/egress_nat_gw/#section)
- [egress_virtual_private_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/egress_virtual_private_gateway/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
egress_gateway_default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
