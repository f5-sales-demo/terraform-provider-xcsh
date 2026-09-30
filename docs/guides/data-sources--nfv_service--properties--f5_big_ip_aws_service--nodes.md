---
page_title: "f5_big_ip_aws_service.nodes"
subcategory: ""
description: "f5_big_ip_aws_service.nodes for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 5448, "body_sha256": "sha256:476bd0a737f4f8b6a1a059c17c8ae1954f167310a6b32ed31fe67382f83c42b6", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "docs/guides/data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.nodes for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# f5_big_ip_aws_service.nodes

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md)
- f5_big_ip_aws_service.nodes

<a id="section"></a>

Type: `"list"`. Computed.

Specify how and where the service nodes are spawned.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [automatic_prefix](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--automatic_prefix.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Computed.

The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the
same Region as your TGW Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--node_name"></a>

### node_name property

Type: `"string"`. Computed.

Node Name will be used to assign as hostname to the service.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [reserved_mgmt_subnet](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--reserved_mgmt_subnet.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--tunnel_prefix"></a>

### tunnel_prefix property

Type: `"string"`. Computed.

Exclusive with \[automatic\_prefix\] Enter IP prefix for the tunnel, it has to be /30.

Upstream description:

Exclusive with \[automatic\_prefix\] Enter IP prefix for the tunnel, it has to be /30.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

## Next pages

- [f5_big_ip_aws_service.nodes.automatic_prefix](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--automatic_prefix.md)
- [f5_big_ip_aws_service.nodes.mgmt_subnet](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet.md)
- [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](data-sources--nfv_service--properties--f5_big_ip_aws_service--nodes--reserved_mgmt_subnet.md)
- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
