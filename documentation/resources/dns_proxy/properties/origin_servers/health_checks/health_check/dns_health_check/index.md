---
page_title: "origin_servers.health_checks.health_check.dns_health_check"
subcategory: ""
description: "DNS health check reports healthy if DNS query is successful and response header and answer matches the given value."
xcsh_docs: {"aliases": ["origin servers health checks health check dns health check", "succeeded", "success", "successful"], "body_bytes": 7941, "body_sha256": "sha256:0a28d700a716d4c20969046695eff347b1bf5158b051a3afeec5fbd57261d98a", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "path": "documentation/resources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0311002023021013-3123212131013012-3232102023122332-1133132222223232-2100231333113223-0323120223013321-3223033220200132-2100333133312221", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [{"anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_response", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.dns_health_check:RequiredObjectAttributes:expected_response,query_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--query_name", "enforcement": "provider-schema", "group": "origin_servers.health_checks.health_check.dns_health_check:RequiredObjectAttributes:expected_response,query_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check"], "schema_version": 1, "sections": [{"aliases": ["origin servers health checks health check dns health check expected rcode"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode", "description": "Expected DNS Response Rcode Type - DNS_RES_RCODE_NOERROR: Rcode NOERROR - DNS_RES_RCODE_ANY: RCODE ANY.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DNS_RES_RCODE_ANY", "DNS_RES_RCODE_NOERROR"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "expected_rcode"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check expected record type"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_record_type", "description": "DNS Response Record Type - DNS_REQUESTED_QUERY_TYPE: Requested Query Type - DNS_RES_RECORD_TYPE_ANY: Any Record Type.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DNS_REQUESTED_QUERY_TYPE", "DNS_RES_RECORD_TYPE_ANY"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "expected_record_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check expected response"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--expected_response", "description": "Specifies an IPv4 or IPv6 address in the answer section of DNS Response.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "expected_response"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check query name"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--query_name", "description": "The query name that the monitor sends a DNS query for.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "query_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check query type"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--query_type", "description": "DNS Query Type - DNS_QTYPE_A: Query Type A - DNS_QTYPE_AAAA: Query Type AAAA.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DNS_QTYPE_A", "DNS_QTYPE_AAAA"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "query_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers health checks health check dns health check reverse", "succeeded", "success", "successful"], "anchor": "schema-origin_servers--health_checks--health_check--dns_health_check--reverse", "description": "Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful receive string match marks the monitored object down instead of up.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check", "reverse"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "DNS health check reports healthy if DNS query is successful and response header and answer matches the given value.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check.dns_health_check

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/)
- [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/health_check/)
- origin_servers.health_checks.health_check.dns_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNS health check reports healthy if DNS query is successful and response header and answer matches
the given value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "query_name")}
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
dns_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode"></a>

### expected_rcode property

Type: `"string"`. Optional.

\[Enum: DNS\_RES\_RCODE\_NOERROR|DNS\_RES\_RCODE\_ANY\] Expected DNS Response Rcode Type -
DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR - DNS\_RES\_RCODE\_ANY: RCODE ANY. Possible values are
\`DNS\_RES\_RCODE\_NOERROR\`, \`DNS\_RES\_RCODE\_ANY\`. Defaults to \`DNS\_RES\_RCODE\_NOERROR\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DNS_RES_RCODE_ANY","DNS_RES_RCODE_NOERROR"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"),
}
```

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

Type: `"string"`. Optional.

\[Enum: DNS\_REQUESTED\_QUERY\_TYPE|DNS\_RES\_RECORD\_TYPE\_ANY\] DNS Response Record Type -
DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type - DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.
Possible values are \`DNS\_REQUESTED\_QUERY\_TYPE\`, \`DNS\_RES\_RECORD\_TYPE\_ANY\`. Defaults to
\`DNS\_REQUESTED\_QUERY\_TYPE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DNS_REQUESTED_QUERY_TYPE","DNS_RES_RECORD_TYPE_ANY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"),
}
```

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

Type: `"string"`. Optional.

Specifies an IPv4 or IPv6 address in the answer section of DNS Response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

Type: `"string"`. Optional.

The query name that the monitor sends a DNS query for.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Type: `"string"`. Optional.

\[Enum: DNS\_QTYPE\_A|DNS\_QTYPE\_AAAA\] DNS Query Type - DNS\_QTYPE\_A: Query Type A -
DNS\_QTYPE\_AAAA: Query Type AAAA. Possible values are \`DNS\_QTYPE\_A\`, \`DNS\_QTYPE\_AAAA\`.
Defaults to \`DNS\_QTYPE\_A\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DNS_QTYPE_A","DNS_QTYPE_AAAA"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"),
}
```

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

Type: `"bool"`. Optional.

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
