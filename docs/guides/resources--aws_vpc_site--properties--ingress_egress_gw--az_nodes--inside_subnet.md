---
page_title: "ingress_egress_gw.az_nodes.inside_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw.az_nodes.inside_subnet for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2981, "body_sha256": "sha256:9827c88f17117662f6cdb4dc088ca75e3f9a8fce9ebb58db3b3a5ebe7f098f47", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.az_nodes.inside_subnet for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.inside_subnet

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes.md)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
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
inside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id"></a>

### existing_subnet_id property

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [subnet_param](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes--inside_subnet--subnet_param.md)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--properties--ingress_egress_gw--az_nodes.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
