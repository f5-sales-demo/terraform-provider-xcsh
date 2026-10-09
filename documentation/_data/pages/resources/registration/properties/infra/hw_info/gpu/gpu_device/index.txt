---
page_title: "infra.hw_info.gpu.gpu_device"
subcategory: ""
description: "List of GPU devices in server."
xcsh_docs: {"aliases": ["infra hw info gpu gpu device"], "body_bytes": 2832, "body_sha256": "sha256:d36f9fa4562ba25df2837947fed0ec8933829166094f8e16266c936ea1f23b2e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "path": "documentation/resources/registration/properties/infra/hw_info/gpu/gpu_device/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1331133021031001-2101301001120200-1031301200113130-3013010302203132-2302103313011011-1030203101033121-3202220230102331-2031132132102023", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "gpu", "gpu_device"], "schema_version": 1, "sections": [{"aliases": ["infra hw info gpu gpu device id"], "anchor": "schema-infra--hw_info--gpu--gpu_device--id", "description": "GPU ID", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info gpu gpu device processes"], "anchor": "schema-infra--hw_info--gpu--gpu_device--processes", "description": "GPU Processes.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device", "processes"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info gpu gpu device product name"], "anchor": "schema-infra--hw_info--gpu--gpu_device--product_name", "description": "GPU Product Name.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device", "product_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/gpu/gpu_device/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of GPU devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["registrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.gpu.gpu_device

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- [infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/gpu/)
- infra.hw_info.gpu.gpu_device

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

GPU devices. List of GPU devices in server.

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
gpu_device {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--gpu--gpu_device--id"></a>

### id property

Type: `"string"`. Optional.

GPU ID. GPU ID

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--gpu--gpu_device--processes"></a>

### processes property

Type: `"string"`. Optional.

Processes. GPU Processes.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--gpu--gpu_device--product_name"></a>

### product_name property

Type: `"string"`. Optional.

Product Name. GPU Product Name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
