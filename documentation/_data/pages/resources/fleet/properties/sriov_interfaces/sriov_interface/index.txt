---
page_title: "sriov_interfaces.sriov_interface"
subcategory: ""
description: "Use custom SR-IOV interfaces Configuration."
xcsh_docs: {"aliases": ["sriov interfaces sriov interface"], "body_bytes": 3255, "body_sha256": "sha256:180a94ebe859827544b683c71574d17d03e82034797387ab8df579363ac5c106", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "parent_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces", "path": "documentation/resources/fleet/properties/sriov_interfaces/sriov_interface/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2200113032031301-2311202302023112-3021303002001132-0201031103201130-3203131323221301-0011022313130111-2212302233122122-2122203210130320", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "type": "requires"}, {"anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["sriov_interfaces", "sriov_interface"], "schema_version": 1, "sections": [{"aliases": ["sriov interfaces sriov interface interface name"], "anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "description": "Name of SR-IOV physical interface.", "document_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "interface_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["sriov interfaces sriov interface number of vfio vfs"], "anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfio_vfs", "description": "Number of virtual functions reserved for VNFs and DPDK-based CNFs.", "document_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "number_of_vfio_vfs"], "syntax": "attribute", "type": "number"}, {"aliases": ["sriov interfaces sriov interface number of vfs"], "anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "description": "Total number of virtual functions.", "document_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "number_of_vfs"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/sriov_interfaces/sriov_interface/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Use custom SR-IOV interfaces Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["fleetCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces.sriov_interface

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [sriov_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/sriov_interfaces/)
- sriov_interfaces.sriov_interface

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Use custom SR-IOV interfaces Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("interface_name",
    "number_of_vfs")}
```

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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

Terraform syntax:

```terraform
sriov_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-sriov_interfaces--sriov_interface--interface_name"></a>

### interface_name property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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

Type: `"number"`. Optional.

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

Type: `"number"`. Optional.

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
