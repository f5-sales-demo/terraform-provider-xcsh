---
page_title: "infra.hw_info.gpu"
subcategory: ""
description: "infra.hw_info.gpu for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 2336, "body_sha256": "sha256:c31fbeed03e0bfe22eb0bf968cf970dce0d85f2a4b8b984a24ae9bfd3cdf725e", "canonical_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu:gpu_device"], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "docs/guides/data-sources--registration--properties--infra--hw_info--gpu.md", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "hw_info", "gpu"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/gpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.hw_info.gpu for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.gpu

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- [infra](data-sources--registration--properties--infra.md)
- [infra.hw_info](data-sources--registration--properties--infra--hw_info.md)
- infra.hw_info.gpu

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-infra--hw_info--gpu--cuda_version"></a>

### cuda_version property

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

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

- [gpu_device](data-sources--registration--properties--infra--hw_info--gpu--gpu_device.md): complete subsection reference.

## Next pages

- [infra.hw_info.gpu.gpu_device](data-sources--registration--properties--infra--hw_info--gpu--gpu_device.md)
- [infra.hw_info](data-sources--registration--properties--infra--hw_info.md)
- [xcsh_registration](../data-sources/registration.md)
