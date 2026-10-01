---
page_title: "f5_big_ip_aws_service"
subcategory: ""
description: "f5_big_ip_aws_service for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 5522, "body_sha256": "sha256:75f7903fef5ef0b41781b23d3079c5223a93bfd97b3887212027dc0cb7381122", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:admin_password", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "parent_id": "xcsh-docs:data-sources:nfv_service:reference", "path": "docs/guides/data-sources--nfv_service--properties--f5_big_ip_aws_service.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- f5_big_ip_aws_service

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Upstream description:

Virtual BIG-IP specification for AWS.

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

- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md#section)
- [palo_alto_fw_service](data-sources--nfv_service--properties--palo_alto_fw_service.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [admin_password](data-sources--nfv_service--properties--f5_big_ip_aws_service--admin_password.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--admin_username"></a>

### admin_username property

Type: `"string"`. Computed.

Admin Username. Admin Username for BIG-IP.

Upstream description:

Admin Username for BIG-IP.

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

- [aws_tgw_site_params](data-sources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params.md): complete subsection reference.

- [endpoint_service](data-sources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md): complete subsection reference.

- [market_place_image](data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md): complete subsection reference.

- [nodes](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--ssh_key"></a>

### ssh_key property

Type: `"string"`. Computed.

Public SSH key for accessing the Big IP nodes.

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

Type: `["map", "string"]`. Computed.

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

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--properties--f5_big_ip_aws_service--admin_password.md)
- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params.md)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md)
- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
