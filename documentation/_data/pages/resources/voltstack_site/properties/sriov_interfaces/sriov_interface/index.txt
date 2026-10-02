---
page_title: "sriov_interfaces.sriov_interface"
subcategory: ""
description: "Use custom SR-IOV interfaces Configuration."
xcsh_docs: {"aliases": ["sriov interfaces sriov interface"], "body_bytes": 3634, "body_sha256": "sha256:dabeaf7fb8af4e3e7b50ac5713504424a20b76b9614180f9866f7fa7eb1068a3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "parent_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces", "path": "documentation/resources/voltstack_site/properties/sriov_interfaces/sriov_interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1211212101221221-1021122103330121-2310112103213310-1123213210212020-3203103101013212-0200213311212023-0223331201133333-0301333023112233", "registry_path": "docs/guides/resources--voltstack_site--reference--group-010.md", "relationships": [{"anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "type": "requires"}, {"anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["sriov_interfaces", "sriov_interface"], "schema_version": 1, "sections": [{"aliases": ["interface name"], "anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "description": "Name of SR-IOV physical interface.", "document_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "interface_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["number of vfio vfs"], "anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfio_vfs", "description": "Number of virtual functions reserved for VNFs and DPDK-based CNFs.", "document_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "number_of_vfio_vfs"], "syntax": "attribute", "type": "number"}, {"aliases": ["number of vfs"], "anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "description": "Total number of virtual functions.", "document_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sriov_interfaces", "sriov_interface", "number_of_vfs"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/sriov_interfaces/sriov_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Use custom SR-IOV interfaces Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces.sriov_interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [sriov_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/sriov_interfaces/)
- sriov_interfaces.sriov_interface

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Use custom SR-IOV interfaces Configuration.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Name of SR-IOV physical interface.

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

Upstream description:

Total number of virtual functions.

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

## Next pages

- [sriov_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/sriov_interfaces/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
