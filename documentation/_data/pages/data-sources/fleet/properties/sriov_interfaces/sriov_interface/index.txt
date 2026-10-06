---
page_title: "sriov_interfaces.sriov_interface"
subcategory: ""
description: "Use custom SR-IOV interfaces Configuration."
xcsh_docs: {"aliases": ["sriov interfaces sriov interface"], "body_bytes": 2934, "body_sha256": "sha256:cae8d3e2e25efa978aa5d8cbbdbda45aac7170123978f15bda123d9733245803", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces:sriov_interface", "parent_id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces", "path": "documentation/data-sources/fleet/properties/sriov_interfaces/sriov_interface/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2223101330132323-2330122331131311-3303222222011201-3323323003030300-0301222013303320-1221121032013123-3311232310000023-0203302130002122", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sriov_interfaces", "sriov_interface"], "schema_version": 1, "sections": [{"aliases": ["sriov interfaces sriov interface interface name"], "anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "description": "Name of SR-IOV physical interface.", "document_id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "interface_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["sriov interfaces sriov interface number of vfio vfs"], "anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfio_vfs", "description": "Number of virtual functions reserved for VNFs and DPDK-based CNFs.", "document_id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "number_of_vfio_vfs"], "syntax": "attribute", "type": "number"}, {"aliases": ["sriov interfaces sriov interface number of vfs"], "anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "description": "Total number of virtual functions.", "document_id": "xcsh-docs:data-sources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "number_of_vfs"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/sriov_interfaces/sriov_interface/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Use custom SR-IOV interfaces Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces.sriov_interface

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [sriov_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/sriov_interfaces/)
- sriov_interfaces.sriov_interface

<a id="section"></a>

Type: `"list"`. Computed.

Use custom SR-IOV interfaces Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-sriov_interfaces--sriov_interface--interface_name"></a>

### interface_name property

Type: `"string"`. Computed.

Name of physical interface. Name of SR-IOV physical interface.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-sriov_interfaces--sriov_interface--number_of_vfio_vfs"></a>

### number_of_vfio_vfs property

Type: `"number"`. Computed.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-sriov_interfaces--sriov_interface--number_of_vfs"></a>

### number_of_vfs property

Type: `"number"`. Computed.

Total number of virtual functions. Total number of virtual functions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
