---
page_title: "tcp_health_check"
subcategory: "Monitoring"
description: "tcp_health_check for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 3394, "body_sha256": "sha256:e8fe1fede8df13d4b03c7c4cabcf374553a7f405ed564418ce089b8da6f86d24", "child_ids": [], "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:properties:tcp_health_check", "parent_id": "xcsh-docs:data-sources:healthcheck:reference", "path": "documentation/data-sources/healthcheck/properties/tcp_health_check/index.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["tcp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/properties/tcp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tcp_health_check for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
