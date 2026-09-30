---
page_title: "f5_big_ip_aws_service"
subcategory: ""
description: "f5_big_ip_aws_service for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 5967, "body_sha256": "sha256:5273d8e8976d4f077e1bb9fe90ce877698a579b2aa68de84301af15b93869308", "canonical_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "docs/guides/resources--nfv_service--properties--f5_big_ip_aws_service.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# f5_big_ip_aws_service

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- f5_big_ip_aws_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Upstream description:

Virtual BIG-IP specification for AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("admin_username",
    "nodes",
    "ssh_key")}
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
  "x-ves-oneof-field-image_choice": "[\"market_place_image\"]",
  "x-ves-oneof-field-site_type_choice": "[\"aws_tgw_site_params\"]"
}
```

OneOf alternatives in this subsection:

- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md#section)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_big_ip_aws_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin_password](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--admin_username"></a>

### admin_username property

Type: `"string"`. Optional.

Admin Username. Admin Username for BIG-IP.

Upstream description:

Admin Username for BIG-IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [aws_tgw_site_params](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params.md): complete subsection reference.

- [endpoint_service](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md): complete subsection reference.

- [market_place_image](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md): complete subsection reference.

- [nodes](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--ssh_key"></a>

### ssh_key property

Type: `"string"`. Optional.

Public SSH key for accessing the Big IP nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-f5_big_ip_aws_service--tags"></a>

### tags property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

## Next pages

- [f5_big_ip_aws_service.admin_password](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password.md)
- [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params.md)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md)
- [f5_big_ip_aws_service.market_place_image](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md)
- [Property reference](resources--nfv_service--reference.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
