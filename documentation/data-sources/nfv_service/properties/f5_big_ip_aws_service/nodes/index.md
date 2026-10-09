---
page_title: "f5_big_ip_aws_service.nodes"
subcategory: ""
description: "Specify how and where the service nodes are spawned."
xcsh_docs: {"aliases": ["f5 big ip aws service nodes"], "body_bytes": 5163, "body_sha256": "sha256:4d6a9b27f652d78184d20d95f757af0020b258292ca67b907633f9b1b6c34759", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3131321100033332-2120130100102110-2101000032232132-2120201123123102-3322330003003013-3102100333000113-3113220113131113-1101122011233323", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service nodes automatic prefix"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "automatic_prefix"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service nodes aws az name"], "anchor": "schema-f5_big_ip_aws_service--nodes--aws_az_name", "description": "The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the same Region as your TGW Site.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service nodes mgmt subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "mgmt_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service nodes node name"], "anchor": "schema-f5_big_ip_aws_service--nodes--node_name", "description": "Node Name will be used to assign as hostname to the service.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "node_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service nodes reserved mgmt subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "reserved_mgmt_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service nodes tunnel prefix"], "anchor": "schema-f5_big_ip_aws_service--nodes--tunnel_prefix", "description": "Exclusive with Enter IP prefix for the tunnel, it has to be /30.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "tunnel_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Specify how and where the service nodes are spawned.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.nodes

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [automatic_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/automatic_prefix/): complete subsection reference.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/): complete subsection reference.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/reserved_mgmt_subnet/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--tunnel_prefix"></a>

### tunnel_prefix property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
