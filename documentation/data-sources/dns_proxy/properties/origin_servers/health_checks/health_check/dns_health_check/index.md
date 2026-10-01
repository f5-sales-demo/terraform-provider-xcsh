---
page_title: "origin_servers.health_checks.health_check.dns_health_check"
subcategory: ""
description: "origin_servers.health_checks.health_check.dns_health_check for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 6849, "body_sha256": "sha256:3e0343218ec7154022a5d553bba138b76e583f7f434a9e9dce701095d2f17ab4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check:dns_health_check", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks:health_check", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/index.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check", "dns_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.health_checks.health_check.dns_health_check for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Expected DNS Response Rcode Type

&#8203;- DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR

&#8203;- DNS\_RES\_RCODE\_ANY: RCODE ANY.

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

Upstream description:

DNS Response Record Type

&#8203;- DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type

&#8203;- DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

DNS Query Type

&#8203;- DNS\_QTYPE\_A: Query Type A

&#8203;- DNS\_QTYPE\_AAAA: Query Type AAAA.

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

Upstream description:

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

## Next pages

- [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
