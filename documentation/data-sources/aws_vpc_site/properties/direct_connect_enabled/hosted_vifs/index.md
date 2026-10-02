---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: "Infrastructure"
description: "AWS Direct Connect Hosted VIF Configuration."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs"], "body_bytes": 2569, "body_sha256": "sha256:b06b8f4b6a7a2fd90360632c156a165faa311bc5dd84bc43066426aff3f1e021", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled", "path": "documentation/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "sections": [{"aliases": ["site registration over direct connect"], "anchor": "section", "description": "CloudLink ADN Network Config.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_direct_connect"], "syntax": "attribute", "type": "object"}, {"aliases": ["site registration over internet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_internet"], "syntax": "attribute", "type": "object"}, {"aliases": ["vif list"], "anchor": "section", "description": "List of Hosted VIF Config.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Direct Connect Hosted VIF Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/)
- direct_connect_enabled.hosted_vifs

<a id="section"></a>

Type: `"single"`. Computed.

AWS Direct Connect Hosted VIF Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

## Direct properties

- [site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/): complete subsection reference.

- [site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/): complete subsection reference.

- [vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/): complete subsection reference.

## Next pages

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/)
- [direct_connect_enabled.hosted_vifs.vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/direct_connect_enabled/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
