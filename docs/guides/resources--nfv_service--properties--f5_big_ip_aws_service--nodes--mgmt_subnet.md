---
page_title: "f5_big_ip_aws_service.nodes.mgmt_subnet"
subcategory: ""
description: "f5_big_ip_aws_service.nodes.mgmt_subnet for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2973, "body_sha256": "sha256:ecce3ae5e795c14e6d93ae34ed8d8b468a5d1696bdd5c619cc6f4cfc4f3b2317", "canonical_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet:subnet_param"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "path": "docs/guides/resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes", "mgmt_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.nodes.mgmt_subnet for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.nodes.mgmt_subnet

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md)
- f5_big_ip_aws_service.nodes.mgmt_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mgmt subnet.

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
mgmt_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-f5_big_ip_aws_service--nodes--mgmt_subnet--existing_subnet_id"></a>

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

- [subnet_param](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param.md)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
