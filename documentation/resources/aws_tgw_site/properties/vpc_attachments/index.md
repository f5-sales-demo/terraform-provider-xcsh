---
page_title: "vpc_attachments"
subcategory: ""
description: "Spoke VPCs to be attached to the AWS TGW Site."
xcsh_docs: {"aliases": ["vpc attachments"], "body_bytes": 951, "body_sha256": "sha256:fc0de424e9615522d9b1f124aadc64988e43a00c4c8c8f689a18b169a9f44e66", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/vpc_attachments/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc_attachments"], "schema_version": 1, "sections": [{"aliases": ["vpc attachments vpc list"], "anchor": "section", "description": "List of VPC attachments to transit gateway.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["vpc_attachments", "vpc_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vpc_attachments/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Spoke VPCs to be attached to the AWS TGW Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
