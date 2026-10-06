---
page_title: "tcp_hex_health_check"
subcategory: ""
description: "Configuration parameter for tcp hex health check."
xcsh_docs: {"aliases": ["tcp hex health check"], "body_bytes": 4490, "body_sha256": "sha256:4d916a4b50b79692b1aade6dabdc37f635d47722defdacf132efa70d549aeaf4", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_hex_health_check", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:reference", "path": "documentation/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2303133011132012-0032300330121222-0130321113203111-3133103302322001-3331022011323230-2300302212313100-3303032312332312-1223301313120331", "registry_path": "docs/guides/data-sources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tcp_hex_health_check"], "schema_version": 1, "sections": [{"aliases": ["tcp hex health check health check port"], "anchor": "schema-tcp_hex_health_check--health_check_port", "description": "Port used for performing health check.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_hex_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_hex_health_check", "health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["tcp hex health check health check secondary port"], "anchor": "schema-tcp_hex_health_check--health_check_secondary_port", "description": "Secondary port used for performing health check. If included, both ports must be healthy for the health check to pass.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_hex_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_hex_health_check", "health_check_secondary_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["tcp hex health check receive"], "anchor": "schema-tcp_hex_health_check--receive", "description": "Hex encoded raw bytes expected in the response.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_hex_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_hex_health_check", "receive"], "syntax": "attribute", "type": "string"}, {"aliases": ["tcp hex health check send"], "anchor": "schema-tcp_hex_health_check--send", "description": "Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check.", "document_id": "xcsh-docs:data-sources:dns_lb_health_check:properties:tcp_hex_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tcp_hex_health_check", "send"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/properties/tcp_hex_health_check/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for tcp hex health check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tcp_hex_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_health_check/properties/)
- tcp_hex_health_check

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp hex health check.

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

<a id="schema-tcp_hex_health_check--health_check_port"></a>

### health_check_port property

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-tcp_hex_health_check--health_check_secondary_port"></a>

### health_check_secondary_port property

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-tcp_hex_health_check--receive"></a>

### receive property

Type: `"string"`. Computed.

Hex encoded raw bytes expected in the response.

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

<a id="schema-tcp_hex_health_check--send"></a>

### send property

Type: `"string"`. Computed.

Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check.

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
