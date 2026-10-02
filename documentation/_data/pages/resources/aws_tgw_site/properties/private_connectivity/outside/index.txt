---
page_title: "private_connectivity.outside"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["private connectivity outside"], "body_bytes": 1267, "body_sha256": "sha256:e55ad634062d0b55a7a708ecb08ed9ce4e9c74befecc1426be78d5f9b9c70a14", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity:outside", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity", "path": "documentation/resources/aws_tgw_site/properties/private_connectivity/outside/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0120301302031320-1033020223303023-1321023321022311-1101000120300312-2200002021233212-2103030102312321-1122032310022332-2312011301133033", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_connectivity", "outside"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/private_connectivity/outside/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity.outside

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/)
- private_connectivity.outside

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
outside = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
