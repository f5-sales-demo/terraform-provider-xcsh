---
page_title: "infra.hw_info.gpu.gpu_device"
subcategory: ""
description: "List of GPU devices in server."
xcsh_docs: {"aliases": ["infra hw info gpu gpu device"], "body_bytes": 2728, "body_sha256": "sha256:9667ab875e0e84fce91b2abb49d064ed8c2ec304203bab5cd9f61fac983deb6c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu:gpu_device", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "path": "documentation/data-sources/registration/properties/infra/hw_info/gpu/gpu_device/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2020012123012313-0100232310221002-3123303132303113-0201101333322121-3312232230123200-2303102232331222-3013233030002123-1203331312133012", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "gpu", "gpu_device"], "schema_version": 1, "sections": [{"aliases": ["infra hw info gpu gpu device id"], "anchor": "schema-infra--hw_info--gpu--gpu_device--id", "description": "GPU ID", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info gpu gpu device processes"], "anchor": "schema-infra--hw_info--gpu--gpu_device--processes", "description": "GPU Processes.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device", "processes"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info gpu gpu device product name"], "anchor": "schema-infra--hw_info--gpu--gpu_device--product_name", "description": "GPU Product Name.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device", "product_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/gpu/gpu_device/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of GPU devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["registrationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.gpu.gpu_device

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/)
- infra.hw_info.gpu.gpu_device

<a id="section"></a>

Type: `"list"`. Computed.

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

## Direct properties

<a id="schema-infra--hw_info--gpu--gpu_device--id"></a>

### id property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
