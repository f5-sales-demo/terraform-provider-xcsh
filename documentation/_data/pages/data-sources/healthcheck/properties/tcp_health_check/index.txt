---
page_title: "tcp_health_check"
subcategory: "Monitoring"
description: "Healthy if TCP connection is successful and response payload matches <expected_response>"
xcsh_docs: {"aliases": ["succeeded", "success", "successful", "tcp health check"], "body_bytes": 3394, "body_sha256": "sha256:728777e47c611f6067282666038a9c408daec9fb3d7e5000e247d9d770ab005f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:properties:tcp_health_check", "parent_id": "xcsh-docs:data-sources:healthcheck:reference", "path": "documentation/data-sources/healthcheck/properties/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1330301203333120-1033112131230233-2331230222230103-3210322303111102-0231333321033120-1333231120301121-1132033312010021-3232320102232110", "registry_path": "docs/guides/data-sources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["tcp health check expected response"], "anchor": "schema-tcp_health_check--expected_response", "description": "Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex encoded payload.", "document_id": "xcsh-docs:data-sources:healthcheck:properties:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["tcp health check send payload"], "anchor": "schema-tcp_health_check--send_payload", "description": "Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the encoding of the payload bytes in the payload. Hex encoded payload.", "document_id": "xcsh-docs:data-sources:healthcheck:properties:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_health_check", "send_payload"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/properties/tcp_health_check/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Healthy if TCP connection is successful and response payload matches <expected_response>", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["healthcheckCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tcp_health_check

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/)
- tcp_health_check

<a id="section"></a>

Type: `"single"`. Computed.

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;.

Upstream description:

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;

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

<a id="schema-tcp_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Computed.

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Upstream description:

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Computed.

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Upstream description:

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/)
