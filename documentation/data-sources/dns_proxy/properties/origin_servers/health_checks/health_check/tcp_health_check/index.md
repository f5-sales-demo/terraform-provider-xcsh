---
page_title: "origin_servers.health_checks.health_check.tcp_health_check"
subcategory: ""
description: "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern."
xcsh_docs: {"aliases": ["login success", "origin servers health checks health check tcp health check", "succeeded", "success", "successful"], "body_bytes": 3547, "body_sha256": "sha256:1f482aecd0a1a21b49aa8c5f1ac4b4fa07f3238fdf7afd6ea947b056adeabf1f", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0210222031302311-3203021130323311-2101200201210030-0301320231303232-1113221323222013-1220230322321332-0301321130133311-2222031032113100", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["expected response"], "anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response", "description": "Specifies a regular expression pattern which will be matched against response payload.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["send payload"], "anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload", "description": "Text string sent in the request.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check", "send_payload"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check.tcp_health_check

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/)
- [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/)
- origin_servers.health_checks.health_check.tcp_health_check

<a id="section"></a>

Type: `"single"`. Computed.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

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

<a id="schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Computed.

Specifies a regular expression pattern which will be matched against response payload.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload"></a>

### send_payload property

Type: `"string"`. Computed.

Send string. Text string sent in the request.

Upstream description:

Text string sent in the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

## Next pages

- [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
