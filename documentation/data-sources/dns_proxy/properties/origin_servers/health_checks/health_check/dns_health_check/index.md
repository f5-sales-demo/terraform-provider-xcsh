---
page_title: "origin_servers.health_checks.health_check.dns_health_check"
subcategory: ""
description: "DNS health check reports healthy if DNS query is successful and response header and answer matches the given value."
xcsh_docs: {"aliases": ["origin servers health checks health check dns health check", "succeeded", "success", "successful"], "body_bytes": 5928, "body_sha256": "sha256:bef03380c60e26191e756066a4953cb18d82d903c31a7f9d5fc0c6a4c4bda274", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-0312002212212310-1033020011102231-3132201333011310-0231131210110000-1233123020130232-2103231311123121-2200303122030210-1110233021320201", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check"], "schema_version": 1, "sections": [{"aliases": ["origin servers health checks health check dns health check expected rcode"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode", "description": "Expected DNS Response Rcode Type - DNS_RES_RCODE_NOERROR: Rcode NOERROR - DNS_RES_RCODE_ANY: RCODE ANY.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "expected_rcode"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check expected record type"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_record_type", "description": "DNS Response Record Type - DNS_REQUESTED_QUERY_TYPE: Requested Query Type - DNS_RES_RECORD_TYPE_ANY: Any Record Type.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "expected_record_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check expected response"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_response", "description": "Specifies an IPv4 or IPv6 address in the answer section of DNS Response.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check query name"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--query_name", "description": "The query name that the monitor sends a DNS query for.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "query_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check query type"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--query_type", "description": "DNS Query Type - DNS_QTYPE_A: Query Type A - DNS_QTYPE_AAAA: Query Type AAAA.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "query_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check reverse", "succeeded", "success", "successful"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--reverse", "description": "Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful receive string match marks the monitored object down instead of up.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "reverse"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "DNS health check reports healthy if DNS query is successful and response header and answer matches the given value.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check.dns_health_check

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/)
- [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/)
- origin_servers.health_checks.health_check.dns_health_check

<a id="section"></a>

Type: `"single"`. Computed.

DNS health check reports healthy if DNS query is successful and response header and answer matches
the given value.

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

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode"></a>

### expected_rcode property

Type: `"string"`. Computed.

\[Enum: DNS\_RES\_RCODE\_NOERROR|DNS\_RES\_RCODE\_ANY\] Expected DNS Response Rcode Type -
DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR - DNS\_RES\_RCODE\_ANY: RCODE ANY. Possible values are
\`DNS\_RES\_RCODE\_NOERROR\`, \`DNS\_RES\_RCODE\_ANY\`. Defaults to \`DNS\_RES\_RCODE\_NOERROR\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_RES_RCODE_NOERROR",
  "enum": [
    "DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--expected_record_type"></a>

### expected_record_type property

Type: `"string"`. Computed.

\[Enum: DNS\_REQUESTED\_QUERY\_TYPE|DNS\_RES\_RECORD\_TYPE\_ANY\] DNS Response Record Type -
DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type - DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.
Possible values are \`DNS\_REQUESTED\_QUERY\_TYPE\`, \`DNS\_RES\_RECORD\_TYPE\_ANY\`. Defaults to
\`DNS\_REQUESTED\_QUERY\_TYPE\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_REQUESTED_QUERY_TYPE",
  "enum": [
    "DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Computed.

Specifies an IPv4 or IPv6 address in the answer section of DNS Response.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--query_name"></a>

### query_name property

Type: `"string"`. Computed.

The query name that the monitor sends a DNS query for.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--query_type"></a>

### query_type property

Type: `"string"`. Computed.

\[Enum: DNS\_QTYPE\_A|DNS\_QTYPE\_AAAA\] DNS Query Type - DNS\_QTYPE\_A: Query Type A -
DNS\_QTYPE\_AAAA: Query Type AAAA. Possible values are \`DNS\_QTYPE\_A\`, \`DNS\_QTYPE\_AAAA\`.
Defaults to \`DNS\_QTYPE\_A\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_QTYPE_A",
  "enum": [
    "DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--reverse"></a>

### reverse property

Type: `"bool"`. Computed.

Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful
receive string match marks the monitored object down instead of up.

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
