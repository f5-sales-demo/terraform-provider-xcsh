---
page_title: "palo_alto_fw_service.service_nodes.nodes"
subcategory: ""
description: "Configuration parameter for nodes"
xcsh_docs: {"aliases": ["palo alto fw service service nodes nodes"], "body_bytes": 5820, "body_sha256": "sha256:7a70fcd092562bfc0dba71eafda22cd95f5868f8e9b46188c2c2f4b84702cb0c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes.nodes:ConflictingListObjectAttributes:mgmt_subnet,reserved_mgmt_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes.nodes:ConflictingListObjectAttributes:mgmt_subnet,reserved_mgmt_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "type": "conflicts"}, {"anchor": "schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes.nodes:RequiredListObjectAttributes:aws_az_name,node_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "type": "requires"}, {"anchor": "schema-palo_alto_fw_service--service_nodes--nodes--node_name", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes.nodes:RequiredListObjectAttributes:aws_az_name,node_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes"], "schema_version": 1, "sections": [{"aliases": ["aws az name"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is one of the AZ for sites.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["mgmt subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes.nodes.mgmt_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes.nodes.mgmt_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["node name"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--node_name", "description": "Node Name will be used to assign as hostname to the service.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "node_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["reserved mgmt subnet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "reserved_mgmt_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for nodes", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Configuration parameter for nodes

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name",
    "node_name"),
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

Provider validators and defaults (from schema source):

```go
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

- [mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/): complete subsection reference.

<a id="schema-palo_alto_fw_service--service_nodes--nodes--node_name"></a>

### node_name property

Type: `"string"`. Optional.

Node Name will be used to assign as hostname to the service.

Provider validators and defaults (from schema source):

```go
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

- [reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/)
- [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
