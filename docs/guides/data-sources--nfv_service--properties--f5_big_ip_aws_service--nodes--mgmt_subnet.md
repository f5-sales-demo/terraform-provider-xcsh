---
page_title: "f5_big_ip_aws_service.nodes.mgmt_subnet"
subcategory: ""
description: "f5_big_ip_aws_service.nodes.mgmt_subnet for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2560, "body_sha256": "sha256:60256de3c9a1d2d19ac68175d4f33d528cb227ee74293dc56a05fafe94435e0b", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet:subnet_param"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "path": "docs/guides/data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes", "mgmt_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.nodes.mgmt_subnet for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.nodes.mgmt_subnet

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes.md)
- f5_big_ip_aws_service.nodes.mgmt_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

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

## Direct properties

<a id="schema-f5_big_ip_aws_service--nodes--mgmt_subnet--existing_subnet_id"></a>

### existing_subnet_id property

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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

- [subnet_param](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param.md)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
