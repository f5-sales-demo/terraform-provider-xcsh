---
page_title: "aws_parameters.az_nodes.outside_subnet"
subcategory: ""
description: "Parameters for AWS subnet."
xcsh_docs: {"aliases": ["aws parameters az nodes outside subnet"], "body_bytes": 2863, "body_sha256": "sha256:646ae5b1e68ef94c254cd9327d92a08fa99e3ec134805cdae25a7fac0a8538c5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0113001211031102-3102103332211013-3221230222101122-1322113203000301-1320331302211332-3023313203312201-3102121331112020-1033102223203322", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.outside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.outside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet:subnet_param", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "az_nodes", "outside_subnet"], "schema_version": 1, "sections": [{"aliases": ["aws parameters az nodes outside subnet existing subnet id"], "anchor": "schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id", "description": "Exclusive with Information about existing subnet ID.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "outside_subnet", "existing_subnet_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters az nodes outside subnet subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet:subnet_param", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--az_nodes--outside_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.outside_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:outside_subnet:subnet_param", "type": "requires"}], "schema_path": ["aws_parameters", "az_nodes", "outside_subnet", "subnet_param"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Parameters for AWS subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.az_nodes.outside_subnet

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [aws_parameters.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/)
- aws_parameters.az_nodes.outside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

Additional upstream details:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id"></a>

### existing_subnet_id property

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/subnet_param/): complete subsection reference.
