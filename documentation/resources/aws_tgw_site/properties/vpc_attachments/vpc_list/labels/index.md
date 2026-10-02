---
page_title: "vpc_attachments.vpc_list.labels"
subcategory: ""
description: "Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall."
xcsh_docs: {"aliases": ["vpc attachments vpc list labels"], "body_bytes": 1432, "body_sha256": "sha256:637b641ead4763c1f01a6f39d2505d6de7574fe2a4ef3c1c3a12ccf4fcacf994", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list", "path": "documentation/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3030120202031300-1103002110210310-2320233211301011-3102302231331323-3201002123312313-2032312122232301-3012212022301011-3313030103300002", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc_attachments", "vpc_list", "labels"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc_attachments.vpc_list.labels

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/)
- [vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/)
- vpc_attachments.vpc_list.labels

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

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
labels {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
