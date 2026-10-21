---
page_title: "create_cloud_hosted.testing"
subcategory: ""
description: "Testing"
xcsh_docs: {"aliases": ["create cloud hosted testing"], "body_bytes": 1895, "body_sha256": "sha256:3ae089dff7bc7b7b74c4e2a2b0bf0e4ff0e81fa0974f02611cb001ebf7871350", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "parent_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "path": "documentation/resources/bot_infrastructure/properties/create_cloud_hosted/testing/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3033321133303122-0331120203303201-2002100100022211-3211122031301311-3321333331333303-0210211002333012-3221003021312121-1220022220231213", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [{"anchor": "schema-create_cloud_hosted--testing--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.testing:RequiredObjectAttributes:region_1", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted", "testing"], "schema_version": 1, "sections": [{"aliases": ["create cloud hosted testing region 1"], "anchor": "schema-create_cloud_hosted--testing--region_1", "description": "This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "testing", "region_1"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/testing/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Testing", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted.testing

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/)
- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/)
- create_cloud_hosted.testing

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Testing

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1")}
```

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
testing {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-create_cloud_hosted--testing--region_1"></a>

### region_1 property

Type: `"string"`. Optional.

This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
