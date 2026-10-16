---
page_title: "tcp_health_check"
subcategory: "Monitoring"
description: "Healthy if TCP connection is successful and response payload matches <expected_response>"
xcsh_docs: {"aliases": ["succeeded", "success", "successful", "tcp health check"], "body_bytes": 3166, "body_sha256": "sha256:8a0c225698a66d0c140244eee94a03dc15ce3a80000bd0a0f2b0481711862b1d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:tcp_health_check", "parent_id": "xcsh-docs:resources:healthcheck:reference", "path": "documentation/resources/healthcheck/properties/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3103201213110121-0000211233303022-2112302331133111-1001213022123321-3021132031330212-3210101201133230-2323313331201022-0113112312002002", "registry_path": "docs/guides/resources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["tcp health check expected response"], "anchor": "schema-tcp_health_check--expected_response", "description": "Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex encoded payload.", "document_id": "xcsh-docs:resources:healthcheck:properties:tcp_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["tcp health check send payload"], "anchor": "schema-tcp_health_check--send_payload", "description": "Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the encoding of the payload bytes in the payload. Hex encoded payload.", "document_id": "xcsh-docs:resources:healthcheck:properties:tcp_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "send_payload"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/tcp_health_check/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Healthy if TCP connection is successful and response payload matches <expected_response>", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["healthcheckCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tcp_health_check

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- tcp_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;.

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
tcp_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tcp_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Optional.

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="schema-tcp_health_check--send_payload"></a>

### send_payload property

Type: `"string"`. Optional.

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```
