---
page_title: "infra.hw_info.memory"
subcategory: ""
description: "infra.hw_info.memory for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 2188, "body_sha256": "sha256:4b1700c5cea22bad5e9f96bbcd5c41cfa3517c8c221a1b6f7a38b41bad068fbe", "canonical_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "child_ids": [], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "docs/guides/data-sources--registration--properties--infra--hw_info--memory.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "hw_info", "memory"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/memory/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.hw_info.memory for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.memory

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- [infra](data-sources--registration--properties--infra.md)
- [infra.hw_info](data-sources--registration--properties--infra--hw_info.md)
- infra.hw_info.memory

<a id="section"></a>

Type: `"single"`. Computed.

Memory Information. Memory information.

Upstream description:

Memory information.

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

<a id="schema-infra--hw_info--memory--size_mb"></a>

### size_mb property

Type: `"number"`. Computed.

RAM. RAM size in MB.

Upstream description:

RAM size in MB.

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

<a id="schema-infra--hw_info--memory--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. RAM data rate in MT/s.

Upstream description:

RAM data rate in MT/s.

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

<a id="schema-infra--hw_info--memory--type"></a>

### type property

Type: `"string"`. Computed.

Type. Type of memory, eg. DDR4.

Upstream description:

Type of memory, eg. DDR4.

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
