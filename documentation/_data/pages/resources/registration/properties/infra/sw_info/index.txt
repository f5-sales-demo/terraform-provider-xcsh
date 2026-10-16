---
page_title: "infra.sw_info"
subcategory: ""
description: "SWInfo holds information about sw version."
xcsh_docs: {"aliases": ["infra sw info"], "body_bytes": 1434, "body_sha256": "sha256:a20066d321624f4cde952a45402eb4ac5abf4cb69786874ed5c61654a5fd1891", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:sw_info", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "documentation/resources/registration/properties/infra/sw_info/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0312321212221203-3210322201111303-0121010122010333-2311320312012111-2333222213001230-0231301313330001-3231102122010310-3030021133012220", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "sw_info"], "schema_version": 1, "sections": [{"aliases": ["infra sw info sw version"], "anchor": "schema-infra--sw_info--sw_version", "description": "SW Version in the site.", "document_id": "xcsh-docs:resources:registration:properties:infra:sw_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "sw_info", "sw_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/sw_info/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "SWInfo holds information about sw version.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["registrationCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.sw_info

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- infra.sw_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SWInfo holds information about sw version.

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
sw_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--sw_info--sw_version"></a>

### sw_version property

Type: `"string"`. Optional.

SW Version. SW Version in the site.

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
