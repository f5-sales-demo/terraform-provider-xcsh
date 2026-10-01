---
page_title: "sriov_interfaces.sriov_interface"
subcategory: ""
description: "sriov_interfaces.sriov_interface for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3314, "body_sha256": "sha256:906d4165057b47f4f3d346a01a94b8823fbfd1b2a0e5969494fe130283796ea6", "canonical_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "parent_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces", "path": "docs/guides/resources--fleet--properties--sriov_interfaces--sriov_interface.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sriov_interfaces", "sriov_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/sriov_interfaces/sriov_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sriov_interfaces.sriov_interface for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces.sriov_interface

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [sriov_interfaces](resources--fleet--properties--sriov_interfaces.md)
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

- [sriov_interfaces](resources--fleet--properties--sriov_interfaces.md)
- [xcsh_fleet](../resources/fleet.md)
