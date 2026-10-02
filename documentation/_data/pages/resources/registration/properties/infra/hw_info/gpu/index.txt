---
page_title: "infra.hw_info.gpu"
subcategory: ""
description: "GPU information on server."
xcsh_docs: {"aliases": ["infra hw info gpu"], "body_bytes": 2826, "body_sha256": "sha256:1647b2f1c18eb1dd27836f7380f5b703f67e42f12532c369f084c4a7905c4d54", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/gpu/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2011333333132333-3300333110123203-0003012331203003-2213312012303211-2201121301032100-0132123323332222-1133303302222001-3030021001011310", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "gpu"], "schema_version": 1, "sections": [{"aliases": ["cuda version"], "anchor": "schema-infra--hw_info--gpu--cuda_version", "description": "GPU Cuda Version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "cuda_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["driver version"], "anchor": "schema-infra--hw_info--gpu--driver_version", "description": "GPU Driver Version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "driver_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["gpu device"], "anchor": "section", "description": "List of GPU devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu:gpu_device", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "gpu", "gpu_device"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/gpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "GPU information on server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

GPU information on server.

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

Upstream description:

GPU Cuda Version.

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

Upstream description:

GPU Driver Version.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [gpu_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/gpu/gpu_device/): complete subsection reference.

## Next pages

- [infra.hw_info.gpu.gpu_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/gpu/gpu_device/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
