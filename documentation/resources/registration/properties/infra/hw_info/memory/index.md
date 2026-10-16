---
page_title: "infra.hw_info.memory"
subcategory: ""
description: "Memory information."
xcsh_docs: {"aliases": ["infra hw info memory"], "body_bytes": 2167, "body_sha256": "sha256:ff8d5dfca9572f2b90936a6305066f43d6e7529ea306817fd0ca1f9427a63cb5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/memory/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3113330200330332-3001000333122232-3332202221130100-2030231103312000-1003212113122233-2201110030130303-0200201103002202-0021320013330100", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "memory"], "schema_version": 1, "sections": [{"aliases": ["infra hw info memory size mb"], "anchor": "schema-infra--hw_info--memory--size_mb", "description": "RAM size in MB.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "memory", "size_mb"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info memory speed"], "anchor": "schema-infra--hw_info--memory--speed", "description": "RAM data rate in MT/s.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "memory", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info memory type"], "anchor": "schema-infra--hw_info--memory--type", "description": "Type of memory, eg. DDR4.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "memory", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/memory/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Memory information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["registrationCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.memory

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- infra.hw_info.memory

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Memory Information. Memory information.

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
memory {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--memory--size_mb"></a>

### size_mb property

Type: `"number"`. Optional.

RAM. RAM size in MB.

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

Type: `"number"`. Optional.

Speed. RAM data rate in MT/s.

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

Type: `"string"`. Optional.

Type. Type of memory, eg. DDR4.

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
