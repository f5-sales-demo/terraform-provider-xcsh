---
page_title: "disable_internet_vip"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable internet vip"], "body_bytes": 1589, "body_sha256": "sha256:6a74e0354e17f19e63920c316aa2ed11282114a3f1ef2b762d7d8f97f461b61a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:disable_internet_vip", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/disable_internet_vip/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1000122022313012-0001101010202330-0013311322231203-1310123033123032-2000220333300012-0323120330133221-0332121010210123-3232302323030130", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_internet_vip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/disable_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_internet_vip

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- disable_internet_vip

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_internet\_vip, enable\_internet\_vip; Default: disable\_internet\_vip\] Enable
this option

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

- [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/disable_internet_vip/#section)
- [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/enable_internet_vip/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_internet_vip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
