---
page_title: "infra.hw_info.kernel"
subcategory: ""
description: "Kernel information."
xcsh_docs: {"aliases": ["infra hw info kernel"], "body_bytes": 2657, "body_sha256": "sha256:c0f5bf6cd260c2475b770abf2974416b0e24d4aefd7cc7c28b72b9d0bc62af4e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/kernel/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3103323320112222-1021320331211233-2330232112001003-3022233110313322-3100023302310032-2112032330023102-0221131110310001-0011212333321203", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "kernel"], "schema_version": 1, "sections": [{"aliases": ["infra hw info kernel architecture"], "anchor": "schema-infra--hw_info--kernel--architecture", "description": "Kernel architecture.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "kernel", "architecture"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info kernel release"], "anchor": "schema-infra--hw_info--kernel--release", "description": "Kernel release.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "kernel", "release"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info kernel version"], "anchor": "schema-infra--hw_info--kernel--version", "description": "Kernel version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "kernel", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/kernel/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Kernel information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["registrationCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.kernel

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- infra.hw_info.kernel

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Kernel. Kernel information.

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
kernel {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--kernel--architecture"></a>

### architecture property

Type: `"string"`. Optional.

Architecture. Kernel architecture.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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

Type: `"string"`. Optional.

Release. Kernel release.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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

Type: `"string"`. Optional.

Version. Kernel version.

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
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
