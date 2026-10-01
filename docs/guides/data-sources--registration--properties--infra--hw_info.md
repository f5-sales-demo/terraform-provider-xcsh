---
page_title: "infra.hw_info"
subcategory: ""
description: "infra.hw_info for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 3878, "body_sha256": "sha256:ad66e8af76a51fade1f04ecf9f902f8d539d75b6d5f79283e29d009125c76245", "canonical_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra:hw_info:bios", "xcsh-docs:data-sources:registration:properties:infra:hw_info:board", "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "xcsh-docs:data-sources:registration:properties:infra:hw_info:product", "xcsh-docs:data-sources:registration:properties:infra:hw_info:storage", "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb"], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "docs/guides/data-sources--registration--properties--infra--hw_info.md", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "hw_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.hw_info for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- [infra](data-sources--registration--properties--infra.md)
- infra.hw_info

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [bios](data-sources--registration--properties--infra--hw_info--bios.md): complete subsection reference.

- [board](data-sources--registration--properties--infra--hw_info--board.md): complete subsection reference.

- [chassis](data-sources--registration--properties--infra--hw_info--chassis.md): complete subsection reference.

- [cpu](data-sources--registration--properties--infra--hw_info--cpu.md): complete subsection reference.

- [gpu](data-sources--registration--properties--infra--hw_info--gpu.md): complete subsection reference.

- [kernel](data-sources--registration--properties--infra--hw_info--kernel.md): complete subsection reference.

- [memory](data-sources--registration--properties--infra--hw_info--memory.md): complete subsection reference.

- [network](data-sources--registration--properties--infra--hw_info--network.md): complete subsection reference.

<a id="schema-infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

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

- [os](data-sources--registration--properties--infra--hw_info--os.md): complete subsection reference.

- [product](data-sources--registration--properties--infra--hw_info--product.md): complete subsection reference.

- [storage](data-sources--registration--properties--infra--hw_info--storage.md): complete subsection reference.

- [usb](data-sources--registration--properties--infra--hw_info--usb.md): complete subsection reference.

## Next pages

- [infra.hw_info.bios](data-sources--registration--properties--infra--hw_info--bios.md)
- [infra.hw_info.board](data-sources--registration--properties--infra--hw_info--board.md)
- [infra.hw_info.chassis](data-sources--registration--properties--infra--hw_info--chassis.md)
- [infra.hw_info.cpu](data-sources--registration--properties--infra--hw_info--cpu.md)
- [infra.hw_info.gpu](data-sources--registration--properties--infra--hw_info--gpu.md)
- [infra.hw_info.kernel](data-sources--registration--properties--infra--hw_info--kernel.md)
- [infra.hw_info.memory](data-sources--registration--properties--infra--hw_info--memory.md)
- [infra.hw_info.network](data-sources--registration--properties--infra--hw_info--network.md)
- [infra.hw_info.os](data-sources--registration--properties--infra--hw_info--os.md)
- [infra.hw_info.product](data-sources--registration--properties--infra--hw_info--product.md)
- [infra.hw_info.storage](data-sources--registration--properties--infra--hw_info--storage.md)
- [infra.hw_info.usb](data-sources--registration--properties--infra--hw_info--usb.md)
- [infra](data-sources--registration--properties--infra.md)
- [xcsh_registration](../data-sources/registration.md)
