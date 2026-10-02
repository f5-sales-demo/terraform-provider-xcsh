---
page_title: "aws_parameters.az_nodes"
subcategory: ""
description: "Only Single AZ or Three AZ(s) nodes are supported currently."
xcsh_docs: {"aliases": ["aws parameters az nodes"], "body_bytes": 3925, "body_sha256": "sha256:c92bd50d23f42a7d199f0c474aba33f80e22251ebf3e93b9ec9b09a469113456", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/az_nodes/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2321230100003213-3132232213200211-1121100121232120-3100002221131133-0101033202211102-2231030013312201-0111221031013030-2232211210231323", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes:ConflictingListObjectAttributes:inside_subnet,reserved_inside_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes:ConflictingListObjectAttributes:inside_subnet,reserved_inside_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "type": "conflicts"}, {"anchor": "schema-aws_parameters--az_nodes--aws_az_name", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes:RequiredListObjectAttributes:aws_az_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "az_nodes"], "schema_version": 1, "sections": [{"aliases": ["aws az name"], "anchor": "schema-aws_parameters--az_nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["inside subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--az_nodes--inside_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.inside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.inside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["aws_parameters", "az_nodes", "inside_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["outside subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.outside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.outside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["aws_parameters", "az_nodes", "outside_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["reserved inside subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "reserved_inside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["workload subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--az_nodes--workload_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.workload_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.workload_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["aws_parameters", "az_nodes", "workload_subnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/az_nodes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Only Single AZ or Three AZ(s) nodes are supported currently.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.az_nodes

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.az_nodes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name"),
  validators.ConflictingListObjectAttributes("inside_subnet",
    "reserved_inside_subnet")}
```

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

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_parameters--az_nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/): complete subsection reference.

- [outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/): complete subsection reference.

- [reserved_inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/reserved_inside_subnet/): complete subsection reference.

- [workload_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/): complete subsection reference.

## Next pages

- [aws_parameters.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/)
- [aws_parameters.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/)
- [aws_parameters.az_nodes.reserved_inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/reserved_inside_subnet/)
- [aws_parameters.az_nodes.workload_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
