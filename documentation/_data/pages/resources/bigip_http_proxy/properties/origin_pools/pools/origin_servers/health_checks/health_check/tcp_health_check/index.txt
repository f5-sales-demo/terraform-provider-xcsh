---
page_title: "origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check"
subcategory: ""
description: "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern."
xcsh_docs: {"aliases": ["origin pools pools origin servers health checks health check tcp health check", "succeeded", "success", "successful"], "body_bytes": 3804, "body_sha256": "sha256:d6c9b44a67664e8d83e64f30848cd18dd3087f0a41dcb14d961988b76f0540a4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0112001111130222-0002130202100220-0202302331311333-3332221001111213-1113010103001132-0031221212001123-0221330031031112-3223013011123101", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["origin pools pools origin servers health checks health check tcp health check expected response"], "anchor": "schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--expected_response", "description": "Specifies a regular expression pattern which will be matched against response payload.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check", "tcp_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin pools pools origin servers health checks health check tcp health check send payload"], "anchor": "schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--send_payload", "description": "Text string sent in the request.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check", "tcp_health_check", "send_payload"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/)
- [origin_pools.pools.origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/)
- origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--send_payload"></a>

### send_payload property

Type: `"string"`. Optional.

Send string. Text string sent in the request.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
