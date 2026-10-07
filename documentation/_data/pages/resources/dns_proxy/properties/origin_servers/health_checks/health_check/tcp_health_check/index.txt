---
page_title: "origin_servers.health_checks.health_check.tcp_health_check"
subcategory: ""
description: "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern."
xcsh_docs: {"aliases": ["origin servers health checks health check tcp health check", "succeeded", "success", "successful"], "body_bytes": 3846, "body_sha256": "sha256:922e873a80d920941ea155cb26a9a65794faefb422e84462b9f9e7f54d54ed91", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "path": "documentation/resources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3103123232330303-0311113021101123-3331030210020302-3003030010121122-0203210221221131-3133212122221102-1323333123103031-1122233122330031", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [{"anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.tcp_health_check:RequiredObjectAttributes:expected_response,send_payload", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.tcp_health_check:RequiredObjectAttributes:expected_response,send_payload", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check"], "schema_version": 1, "sections": [{"aliases": ["origin servers health checks health check tcp health check expected response"], "anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response", "description": "Specifies a regular expression pattern which will be matched against response payload.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check tcp health check send payload"], "anchor": "schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload", "description": "Text string sent in the request.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:tcp_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "tcp_health_check", "send_payload"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Monitor reports healthy status if UDP connection is successful and response payload matches expected response pattern.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check.tcp_health_check

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/)
- [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/)
- origin_servers.health_checks.health_check.tcp_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "send_payload")}
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Optional.

Specifies a regular expression pattern which will be matched against response payload.

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
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Type: `"string"`. Optional.

Send string. Text string sent in the request.

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
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
