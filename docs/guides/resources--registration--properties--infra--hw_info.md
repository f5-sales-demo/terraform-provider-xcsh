---
page_title: "infra.hw_info"
subcategory: ""
description: "infra.hw_info for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 4036, "body_sha256": "sha256:d136c61eeab69396932fee2b83cccc85907785da6a33f209c8b0b0a62240d781", "canonical_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "child_ids": ["xcsh-docs:resources:registration:properties:infra:hw_info:bios", "xcsh-docs:resources:registration:properties:infra:hw_info:board", "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "xcsh-docs:resources:registration:properties:infra:hw_info:cpu", "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "xcsh-docs:resources:registration:properties:infra:hw_info:network", "xcsh-docs:resources:registration:properties:infra:hw_info:os", "xcsh-docs:resources:registration:properties:infra:hw_info:product", "xcsh-docs:resources:registration:properties:infra:hw_info:storage", "xcsh-docs:resources:registration:properties:infra:hw_info:usb"], "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "docs/guides/resources--registration--properties--infra--hw_info.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "hw_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.hw_info for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info

Breadcrumbs:

- [xcsh_registration](../resources/registration.md)
- [Property reference](resources--registration--reference.md)
- [infra](resources--registration--properties--infra.md)
- infra.hw_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

OsInfo holds information about host OS and HW.

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

Terraform syntax:

```terraform
hw_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bios](resources--registration--properties--infra--hw_info--bios.md): complete subsection reference.

- [board](resources--registration--properties--infra--hw_info--board.md): complete subsection reference.

- [chassis](resources--registration--properties--infra--hw_info--chassis.md): complete subsection reference.

- [cpu](resources--registration--properties--infra--hw_info--cpu.md): complete subsection reference.

- [gpu](resources--registration--properties--infra--hw_info--gpu.md): complete subsection reference.

- [kernel](resources--registration--properties--infra--hw_info--kernel.md): complete subsection reference.

- [memory](resources--registration--properties--infra--hw_info--memory.md): complete subsection reference.

- [network](resources--registration--properties--infra--hw_info--network.md): complete subsection reference.

<a id="schema-infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Optional.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0"
  }
}
```

- [os](resources--registration--properties--infra--hw_info--os.md): complete subsection reference.

- [product](resources--registration--properties--infra--hw_info--product.md): complete subsection reference.

- [storage](resources--registration--properties--infra--hw_info--storage.md): complete subsection reference.

- [usb](resources--registration--properties--infra--hw_info--usb.md): complete subsection reference.

## Next pages

- [infra.hw_info.bios](resources--registration--properties--infra--hw_info--bios.md)
- [infra.hw_info.board](resources--registration--properties--infra--hw_info--board.md)
- [infra.hw_info.chassis](resources--registration--properties--infra--hw_info--chassis.md)
- [infra.hw_info.cpu](resources--registration--properties--infra--hw_info--cpu.md)
- [infra.hw_info.gpu](resources--registration--properties--infra--hw_info--gpu.md)
- [infra.hw_info.kernel](resources--registration--properties--infra--hw_info--kernel.md)
- [infra.hw_info.memory](resources--registration--properties--infra--hw_info--memory.md)
- [infra.hw_info.network](resources--registration--properties--infra--hw_info--network.md)
- [infra.hw_info.os](resources--registration--properties--infra--hw_info--os.md)
- [infra.hw_info.product](resources--registration--properties--infra--hw_info--product.md)
- [infra.hw_info.storage](resources--registration--properties--infra--hw_info--storage.md)
- [infra.hw_info.usb](resources--registration--properties--infra--hw_info--usb.md)
- [infra](resources--registration--properties--infra.md)
- [xcsh_registration](../resources/registration.md)
