---
page_title: "palo_alto_fw_service.service_nodes.nodes"
subcategory: ""
description: "Configuration parameter for nodes"
xcsh_docs: {"aliases": ["palo alto fw service service nodes nodes"], "body_bytes": 4521, "body_sha256": "sha256:6c737e56f58145806fca5ba52c09d470103db0b57100c8c60626425eb68f2b1f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service service nodes nodes aws az name"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is one of the AZ for sites.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["palo alto fw service service nodes nodes mgmt subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["palo alto fw service service nodes nodes node name"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--node_name", "description": "Node Name will be used to assign as hostname to the service.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "node_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["palo alto fw service service nodes nodes reserved mgmt subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "reserved_mgmt_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Configuration parameter for nodes", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes.nodes

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- palo_alto_fw_service.service_nodes.nodes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Palo Alto Networks AZ Nodes. Configuration parameter for nodes

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

<a id="schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is
one of the AZ for sites.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

- [mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/): complete subsection reference.

<a id="schema-palo_alto_fw_service--service_nodes--nodes--node_name"></a>

### node_name property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

- [reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/): complete subsection reference.
