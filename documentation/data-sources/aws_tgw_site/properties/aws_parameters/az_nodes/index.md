---
page_title: "aws_parameters.az_nodes"
subcategory: ""
description: "Only Single AZ or Three AZ(s) nodes are supported currently."
xcsh_docs: {"aliases": ["aws parameters az nodes"], "body_bytes": 3602, "body_sha256": "sha256:b957b1c5b4fb74f49c9932b4d12a2c68058368705ac177ac1d710c0839f71f29", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "path": "documentation/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "az_nodes"], "schema_version": 1, "sections": [{"aliases": ["aws parameters az nodes aws az name"], "anchor": "schema-aws_parameters--az_nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters az nodes inside subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "inside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters az nodes outside subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "outside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters az nodes reserved inside subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "reserved_inside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters az nodes workload subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "workload_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Only Single AZ or Three AZ(s) nodes are supported currently.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.az_nodes

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.az_nodes

<a id="section"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

## Direct properties

<a id="schema-aws_parameters--az_nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/): complete subsection reference.

- [outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/): complete subsection reference.

- [reserved_inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/reserved_inside_subnet/): complete subsection reference.

- [workload_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/): complete subsection reference.

## Next pages

- [aws_parameters.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/)
- [aws_parameters.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/)
- [aws_parameters.az_nodes.reserved_inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/reserved_inside_subnet/)
- [aws_parameters.az_nodes.workload_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
