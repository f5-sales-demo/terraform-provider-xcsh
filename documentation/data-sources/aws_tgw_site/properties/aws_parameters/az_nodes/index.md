---
page_title: "aws_parameters.az_nodes"
subcategory: ""
description: "Only Single AZ or Three AZ(s) nodes are supported currently."
xcsh_docs: {"aliases": ["aws parameters az nodes"], "body_bytes": 2627, "body_sha256": "sha256:ae1430942cb40c2384bf5f02e76933dc31f8e355d6e2fbd0362f63a539ca60ed", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "path": "documentation/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "az_nodes"], "schema_version": 1, "sections": [{"aliases": ["aws parameters az nodes aws az name"], "anchor": "schema-aws_parameters--az_nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters az nodes inside subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "inside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters az nodes outside subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "outside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters az nodes reserved inside subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "reserved_inside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters az nodes workload subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "workload_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/az_nodes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Only Single AZ or Three AZ(s) nodes are supported currently.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
