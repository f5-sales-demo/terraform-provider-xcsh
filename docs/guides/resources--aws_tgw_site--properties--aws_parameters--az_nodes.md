---
page_title: "aws_parameters.az_nodes"
subcategory: ""
description: "aws_parameters.az_nodes for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 3276, "body_sha256": "sha256:e342392fffaae30d60b14afb474bd50b2d7aa414ae8a7e9e651215baa73336d4", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:workload_subnet"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "docs/guides/resources--aws_tgw_site--properties--aws_parameters--az_nodes.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_parameters", "az_nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/az_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.az_nodes for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.az_nodes

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
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

- [inside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet.md): complete subsection reference.

- [outside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet.md): complete subsection reference.

- [reserved_inside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--reserved_inside_subnet.md): complete subsection reference.

- [workload_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet.md): complete subsection reference.

## Next pages

- [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--inside_subnet.md)
- [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--outside_subnet.md)
- [aws_parameters.az_nodes.reserved_inside_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--reserved_inside_subnet.md)
- [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--properties--aws_parameters--az_nodes--workload_subnet.md)
- [aws_parameters](resources--aws_tgw_site--properties--aws_parameters.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
