---
page_title: "infra.hw_info.kernel"
subcategory: ""
description: "infra.hw_info.kernel for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 2666, "body_sha256": "sha256:43e5a4665cefd6d15f0841f4d2fdd5e66bb4a1ebc2ad3c55e7e403c8ac309856", "canonical_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "child_ids": [], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "docs/guides/data-sources--registration--properties--infra--hw_info--kernel.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "hw_info", "kernel"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/kernel/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.hw_info.kernel for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.kernel

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- [infra](data-sources--registration--properties--infra.md)
- [infra.hw_info](data-sources--registration--properties--infra--hw_info.md)
- infra.hw_info.kernel

<a id="section"></a>

Type: `"single"`. Computed.

Kernel. Kernel information.

Upstream description:

Kernel information.

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

<a id="schema-infra--hw_info--kernel--architecture"></a>

### architecture property

Type: `"string"`. Computed.

Architecture. Kernel architecture.

Upstream description:

Kernel architecture.

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

<a id="schema-infra--hw_info--kernel--release"></a>

### release property

Type: `"string"`. Computed.

Release. Kernel release.

Upstream description:

Kernel release.

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

<a id="schema-infra--hw_info--kernel--version"></a>

### version property

Type: `"string"`. Computed.

Version. Kernel version.

Upstream description:

Kernel version.

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

## Next pages

- [infra.hw_info](data-sources--registration--properties--infra--hw_info.md)
- [xcsh_registration](../data-sources/registration.md)
