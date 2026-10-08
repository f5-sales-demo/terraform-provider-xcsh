---
page_title: "infra.hw_info.gpu"
subcategory: ""
description: "GPU information on server."
xcsh_docs: {"aliases": ["infra hw info gpu"], "body_bytes": 2292, "body_sha256": "sha256:0157c97c9020e5a2079e91ec56fe067e06fb16921a7715e1149750cbf541c04f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/gpu/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2011333333132333-3300333110123203-0003012331203003-2213312012303211-2201121301032100-0132123323332222-1133303302222001-3030021001011310", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "gpu"], "schema_version": 1, "sections": [{"aliases": ["infra hw info gpu cuda version"], "anchor": "schema-infra--hw_info--gpu--cuda_version", "description": "GPU Cuda Version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "cuda_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info gpu driver version"], "anchor": "schema-infra--hw_info--gpu--driver_version", "description": "GPU Driver Version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "driver_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info gpu gpu device"], "anchor": "section", "description": "List of GPU devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/gpu/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "GPU information on server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["registrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.gpu

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- infra.hw_info.gpu

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

GPU. GPU information on server.

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
gpu {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--gpu--cuda_version"></a>

### cuda_version property

Type: `"string"`. Optional.

Cuda Version. GPU Cuda Version.

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

<a id="schema-infra--hw_info--gpu--driver_version"></a>

### driver_version property

Type: `"string"`. Optional.

Driver Version. GPU Driver Version.

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

- [gpu_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/gpu/gpu_device/): complete subsection reference.
