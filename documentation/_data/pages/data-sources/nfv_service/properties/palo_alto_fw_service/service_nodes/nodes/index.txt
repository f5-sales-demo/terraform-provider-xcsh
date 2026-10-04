---
page_title: "palo_alto_fw_service.service_nodes.nodes"
subcategory: ""
description: "Configuration parameter for nodes"
xcsh_docs: {"aliases": ["palo alto fw service service nodes nodes"], "body_bytes": 5197, "body_sha256": "sha256:ff78c78fd7ebe0b05a8a2508f731d803dc3b9ea4e6cef645bae91aeca808df12", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0032213330302312-3301201210322002-3313201212101103-0313133031321132-3030013231120110-3303023132223202-1220002233012213-1230213312103102", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service service nodes nodes aws az name"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is one of the AZ for sites.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["palo alto fw service service nodes nodes mgmt subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["palo alto fw service service nodes nodes node name"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--node_name", "description": "Node Name will be used to assign as hostname to the service.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "node_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["palo alto fw service service nodes nodes reserved mgmt subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "reserved_mgmt_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration parameter for nodes", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes.nodes

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- palo_alto_fw_service.service_nodes.nodes

<a id="section"></a>

Type: `"list"`. Computed.

Palo Alto Networks AZ Nodes. Configuration parameter for nodes

Upstream description:

Configuration parameter for nodes

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/): complete subsection reference.

<a id="schema-palo_alto_fw_service--service_nodes--nodes--node_name"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/)
- [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
