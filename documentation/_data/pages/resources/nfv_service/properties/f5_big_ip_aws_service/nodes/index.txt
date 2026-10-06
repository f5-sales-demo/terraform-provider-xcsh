---
page_title: "f5_big_ip_aws_service.nodes"
subcategory: ""
description: "Specify how and where the service nodes are spawned."
xcsh_docs: {"aliases": ["f5 big ip aws service nodes"], "body_bytes": 5973, "body_sha256": "sha256:92cc466da38d9f39161b4ccc0f8fe9e480e671289b3d5c78856fc78c621e87b7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0303213133011232-0201322132022131-1321023330330310-0032213003120332-3011101310033312-2030333102020323-2111032131310203-3332013322121201", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--nodes--tunnel_prefix", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:automatic_prefix,tunnel_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:automatic_prefix,tunnel_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:mgmt_subnet,reserved_mgmt_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:mgmt_subnet,reserved_mgmt_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet", "type": "conflicts"}, {"anchor": "schema-f5_big_ip_aws_service--nodes--aws_az_name", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:RequiredListObjectAttributes:aws_az_name,node_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "requires"}, {"anchor": "schema-f5_big_ip_aws_service--nodes--node_name", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:RequiredListObjectAttributes:aws_az_name,node_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service nodes automatic prefix"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "automatic_prefix"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service nodes aws az name"], "anchor": "schema-f5_big_ip_aws_service--nodes--aws_az_name", "description": "The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the same Region as your TGW Site.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service nodes mgmt subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--nodes--mgmt_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes.mgmt_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes.mgmt_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["f5_big_ip_aws_service", "nodes", "mgmt_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service nodes node name"], "anchor": "schema-f5_big_ip_aws_service--nodes--node_name", "description": "Node Name will be used to assign as hostname to the service.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "node_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service nodes reserved mgmt subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "reserved_mgmt_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service nodes tunnel prefix"], "anchor": "schema-f5_big_ip_aws_service--nodes--tunnel_prefix", "description": "Exclusive with Enter IP prefix for the tunnel, it has to be /30.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "tunnel_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Specify how and where the service nodes are spawned.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.nodes

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- f5_big_ip_aws_service.nodes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Specify how and where the service nodes are spawned.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name",
    "node_name"),
  validators.ConflictingListObjectAttributes("automatic_prefix",
    "tunnel_prefix"),
  validators.ConflictingListObjectAttributes("mgmt_subnet",
    "reserved_mgmt_subnet")}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
nodes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [automatic_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/automatic_prefix/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Optional.

The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the
same Region as your TGW Site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--node_name"></a>

### node_name property

Type: `"string"`. Optional.

Node Name will be used to assign as hostname to the service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/reserved_mgmt_subnet/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--nodes--tunnel_prefix"></a>

### tunnel_prefix property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
