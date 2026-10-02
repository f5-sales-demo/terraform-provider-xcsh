---
page_title: "aws_provider"
subcategory: ""
description: "Cloud Connect with AWS."
xcsh_docs: {"aliases": ["aws provider"], "body_bytes": 1853, "body_sha256": "sha256:938eded216b46f03277e3d767c3a4bd4a7cc03ba449f07d0ca0cfa4b6111a1b4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider", "parent_id": "xcsh-docs:resources:cloud_connect:reference", "path": "documentation/resources/cloud_connect/properties/aws_provider/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider"], "schema_version": 1, "sections": [{"aliases": ["aws tgw site"], "anchor": "section", "description": "Cloud Connect AWS TGW Site Type.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Cloud Connect with AWS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- aws_provider

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_provider, azure\_vnet\_site\] Configuration parameter for aws provider.

Upstream description:

Cloud Connect with AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_type": "[\"aws_tgw_site\"]"
}
```

OneOf alternatives in this subsection:

- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/#section)
- [azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/azure_vnet_site/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

## Direct properties

- [aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
