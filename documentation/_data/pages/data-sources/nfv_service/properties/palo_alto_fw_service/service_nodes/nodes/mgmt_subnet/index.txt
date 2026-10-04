---
page_title: "palo_alto_fw_service.service_nodes.nodes.mgmt_subnet"
subcategory: ""
description: "Parameters for AWS subnet."
xcsh_docs: {"aliases": ["palo alto fw service service nodes nodes mgmt subnet"], "body_bytes": 3259, "body_sha256": "sha256:599e2c8eb3d80fa89ef343b503cefd6aa18b0be5b0fd405d030c9e2d53aa8cf2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0321303313302233-1320211102133320-0210022301030202-0110133100031201-0021112121011033-2112101021033213-2012103021120332-3302201223130130", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service service nodes nodes mgmt subnet existing subnet id"], "anchor": "schema-palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--existing_subnet_id", "description": "Exclusive with Information about existing subnet ID.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet", "existing_subnet_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["palo alto fw service service nodes nodes mgmt subnet subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet", "subnet_param"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Parameters for AWS subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes.nodes.mgmt_subnet

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- [palo_alto_fw_service.service_nodes.nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet

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

<a id="schema-palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--existing_subnet_id"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/subnet_param/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/subnet_param/)
- [palo_alto_fw_service.service_nodes.nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
