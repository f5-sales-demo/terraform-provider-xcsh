---
page_title: "voltstack_cluster.az_nodes.local_subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster.az_nodes.local_subnet for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2874, "body_sha256": "sha256:8ba5051aca4ec3d0f097bcd1cd207a4a359d3bcc888d1e56fe2de9bb337fdc51", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:az_nodes:local_subnet", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:az_nodes", "path": "docs/guides/resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/voltstack_cluster/az_nodes/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.az_nodes.local_subnet for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.az_nodes.local_subnet

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes.md)
- voltstack_cluster.az_nodes.local_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

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
local_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster--az_nodes--local_subnet--existing_subnet_id"></a>

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

- [subnet_param](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes--local_subnet--subnet_param.md)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--properties--voltstack_cluster--az_nodes.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
