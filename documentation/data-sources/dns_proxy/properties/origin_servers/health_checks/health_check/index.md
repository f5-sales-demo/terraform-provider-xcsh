---
page_title: "origin_servers.health_checks.health_check"
subcategory: ""
description: "List of Health Checks."
xcsh_docs: {"aliases": ["origin servers health checks health check"], "body_bytes": 2405, "body_sha256": "sha256:60781284384c9afde679cc4a6b172136383ec7e8a49a1a82f4503e69194568e4", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "data-sources", "registry_anchor": "canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check"], "schema_version": 1, "sections": [{"aliases": ["origin servers health checks health check dns health check", "succeeded", "success", "successful"], "anchor": "section", "description": "DNS health check reports healthy if DNS query is successful and response header and answer matches the given value.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers health checks health check icmp health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "icmp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers health checks health check tcp health check", "succeeded", "success", "successful"], "anchor": "section", "description": "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "List of Health Checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/)
- origin_servers.health_checks.health_check

<a id="section"></a>

Type: `"list"`. Computed.

List of Health Checks. List of Health Checks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-13T13:24:52+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [dns_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/): complete subsection reference.

- [icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/icmp_health_check/): complete subsection reference.

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/): complete subsection reference.
