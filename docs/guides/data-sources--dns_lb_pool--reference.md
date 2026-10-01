---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 18041, "body_sha256": "sha256:e6752c3ca4d78c825df41ee7bc5f9be309a9e25105cf3c291bccecb63090cfb0", "canonical_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool", "xcsh-docs:data-sources:dns_lb_pool:properties:mx_pool", "xcsh-docs:data-sources:dns_lb_pool:properties:srv_pool", "xcsh-docs:data-sources:dns_lb_pool:properties:use_rrset_ttl"], "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:reference", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:fundamentals", "path": "docs/guides/data-sources--dns_lb_pool--reference.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
- Property reference

## Direct properties

- [a_pool](data-sources--dns_lb_pool--properties--a_pool.md): complete subsection reference.

- [aaaa_pool](data-sources--dns_lb_pool--properties--aaaa_pool.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [cname_pool](data-sources--dns_lb_pool--properties--cname_pool.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the DNSLBPool.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-load_balancing_mode"></a>

### load_balancing_mode property

Type: `"string"`. Computed.

\[Enum: ROUND\_ROBIN|RATIO\_MEMBER|STATIC\_PERSIST|PRIORITY\] - ROUND\_ROBIN: Round-Robin Round
Robin will ensure random equal distribution of requests among all pool members in a pool. -
RATIO\_MEMBER: Ratio-Member Ratio-Member performs load balancing of requests across the pool members
based on the ratio assigned to each pool member - STATIC\_PERSIST.. Possible values are
\`ROUND\_ROBIN\`, \`RATIO\_MEMBER\`, \`STATIC\_PERSIST\`, \`PRIORITY\`. Defaults to
\`ROUND\_ROBIN\`.

Upstream description:

&#8203;- ROUND\_ROBIN: Round-Robin

Round Robin will ensure random equal distribution of requests among all pool members in a pool.
&#8203;- RATIO\_MEMBER: Ratio-Member

Ratio-Member performs load balancing of requests across the pool members based on the ratio assigned
to each pool member &#8203;- STATIC\_PERSIST: Static-Persist

The Static Persist load balancing method uses the persist mask, with the source IP address of the
Local Domain Name Server (LDNS), in a deterministic algorithm to send requests to a specific pool
member. If the DNS resolver passes ECS (EDNS-Client-Subnet) information, then a hash of it will be
used, to send the client to the same pool member &#8203;- PRIORITY: Priority

The Priority load balancing method returns all available endpoints in a pool with the highest
priority. Pool Members have a priority value, starting from zero, where a lower value means a higher
priority.

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [mx_pool](data-sources--dns_lb_pool--properties--mx_pool.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the DNSLBPool.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace where the DNSLBPool exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [srv_pool](data-sources--dns_lb_pool--properties--srv_pool.md): complete subsection reference.

<a id="schema-ttl"></a>

### ttl property

Type: `"number"`. Computed.

\[OneOf: ttl, use\_rrset\_ttl\] Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default
&#8203;30) for responses from this pool.

Upstream description:

Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default 30) for responses from this pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

OneOf alternatives in this subsection:

- [ttl](data-sources--dns_lb_pool--reference.md#schema-ttl)
- [use_rrset_ttl](data-sources--dns_lb_pool--properties--use_rrset_ttl.md#section)

Select alternatives according to the provider validators above.

- [use_rrset_ttl](data-sources--dns_lb_pool--properties--use_rrset_ttl.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `a_pool` | [a_pool](data-sources--dns_lb_pool--properties--a_pool.md#section) |
| `a_pool.disable_health_check` | [a_pool.disable_health_check](data-sources--dns_lb_pool--properties--a_pool--disable_health_check.md#section) |
| `a_pool.health_check` | [a_pool.health_check](data-sources--dns_lb_pool--properties--a_pool--health_check.md#section) |
| `a_pool.health_check.name` | [a_pool.health_check.name](data-sources--dns_lb_pool--properties--a_pool--health_check.md#schema-a_pool--health_check--name) |
| `a_pool.health_check.namespace` | [a_pool.health_check.namespace](data-sources--dns_lb_pool--properties--a_pool--health_check.md#schema-a_pool--health_check--namespace) |
| `a_pool.health_check.tenant` | [a_pool.health_check.tenant](data-sources--dns_lb_pool--properties--a_pool--health_check.md#schema-a_pool--health_check--tenant) |
| `a_pool.max_answers` | [a_pool.max_answers](data-sources--dns_lb_pool--properties--a_pool.md#schema-a_pool--max_answers) |
| `a_pool.members` | [a_pool.members](data-sources--dns_lb_pool--properties--a_pool--members.md#section) |
| `a_pool.members.disable_spec` | [a_pool.members.disable_spec](data-sources--dns_lb_pool--properties--a_pool--members.md#schema-a_pool--members--disable_spec) |
| `a_pool.members.ip_endpoint` | [a_pool.members.ip_endpoint](data-sources--dns_lb_pool--properties--a_pool--members.md#schema-a_pool--members--ip_endpoint) |
| `a_pool.members.name` | [a_pool.members.name](data-sources--dns_lb_pool--properties--a_pool--members.md#schema-a_pool--members--name) |
| `a_pool.members.priority` | [a_pool.members.priority](data-sources--dns_lb_pool--properties--a_pool--members.md#schema-a_pool--members--priority) |
| `a_pool.members.ratio` | [a_pool.members.ratio](data-sources--dns_lb_pool--properties--a_pool--members.md#schema-a_pool--members--ratio) |
| `aaaa_pool` | [aaaa_pool](data-sources--dns_lb_pool--properties--aaaa_pool.md#section) |
| `aaaa_pool.max_answers` | [aaaa_pool.max_answers](data-sources--dns_lb_pool--properties--aaaa_pool.md#schema-aaaa_pool--max_answers) |
| `aaaa_pool.members` | [aaaa_pool.members](data-sources--dns_lb_pool--properties--aaaa_pool--members.md#section) |
| `aaaa_pool.members.disable_spec` | [aaaa_pool.members.disable_spec](data-sources--dns_lb_pool--properties--aaaa_pool--members.md#schema-aaaa_pool--members--disable_spec) |
| `aaaa_pool.members.ip_endpoint` | [aaaa_pool.members.ip_endpoint](data-sources--dns_lb_pool--properties--aaaa_pool--members.md#schema-aaaa_pool--members--ip_endpoint) |
| `aaaa_pool.members.name` | [aaaa_pool.members.name](data-sources--dns_lb_pool--properties--aaaa_pool--members.md#schema-aaaa_pool--members--name) |
| `aaaa_pool.members.priority` | [aaaa_pool.members.priority](data-sources--dns_lb_pool--properties--aaaa_pool--members.md#schema-aaaa_pool--members--priority) |
| `aaaa_pool.members.ratio` | [aaaa_pool.members.ratio](data-sources--dns_lb_pool--properties--aaaa_pool--members.md#schema-aaaa_pool--members--ratio) |
| `annotations` | [annotations](data-sources--dns_lb_pool--reference.md#schema-annotations) |
| `cname_pool` | [cname_pool](data-sources--dns_lb_pool--properties--cname_pool.md#section) |
| `cname_pool.disable_health_check` | [cname_pool.disable_health_check](data-sources--dns_lb_pool--properties--cname_pool--disable_health_check.md#section) |
| `cname_pool.health_check` | [cname_pool.health_check](data-sources--dns_lb_pool--properties--cname_pool--health_check.md#section) |
| `cname_pool.health_check.name` | [cname_pool.health_check.name](data-sources--dns_lb_pool--properties--cname_pool--health_check.md#schema-cname_pool--health_check--name) |
| `cname_pool.health_check.namespace` | [cname_pool.health_check.namespace](data-sources--dns_lb_pool--properties--cname_pool--health_check.md#schema-cname_pool--health_check--namespace) |
| `cname_pool.health_check.tenant` | [cname_pool.health_check.tenant](data-sources--dns_lb_pool--properties--cname_pool--health_check.md#schema-cname_pool--health_check--tenant) |
| `cname_pool.members` | [cname_pool.members](data-sources--dns_lb_pool--properties--cname_pool--members.md#section) |
| `cname_pool.members.domain` | [cname_pool.members.domain](data-sources--dns_lb_pool--properties--cname_pool--members.md#schema-cname_pool--members--domain) |
| `cname_pool.members.final_translation` | [cname_pool.members.final_translation](data-sources--dns_lb_pool--properties--cname_pool--members.md#schema-cname_pool--members--final_translation) |
| `cname_pool.members.name` | [cname_pool.members.name](data-sources--dns_lb_pool--properties--cname_pool--members.md#schema-cname_pool--members--name) |
| `cname_pool.members.priority` | [cname_pool.members.priority](data-sources--dns_lb_pool--properties--cname_pool--members.md#schema-cname_pool--members--priority) |
| `cname_pool.members.ratio` | [cname_pool.members.ratio](data-sources--dns_lb_pool--properties--cname_pool--members.md#schema-cname_pool--members--ratio) |
| `description` | [description](data-sources--dns_lb_pool--reference.md#schema-description) |
| `id` | [id](data-sources--dns_lb_pool--reference.md#schema-id) |
| `labels` | [labels](data-sources--dns_lb_pool--reference.md#schema-labels) |
| `load_balancing_mode` | [load_balancing_mode](data-sources--dns_lb_pool--reference.md#schema-load_balancing_mode) |
| `mx_pool` | [mx_pool](data-sources--dns_lb_pool--properties--mx_pool.md#section) |
| `mx_pool.max_answers` | [mx_pool.max_answers](data-sources--dns_lb_pool--properties--mx_pool.md#schema-mx_pool--max_answers) |
| `mx_pool.members` | [mx_pool.members](data-sources--dns_lb_pool--properties--mx_pool--members.md#section) |
| `mx_pool.members.domain` | [mx_pool.members.domain](data-sources--dns_lb_pool--properties--mx_pool--members.md#schema-mx_pool--members--domain) |
| `mx_pool.members.name` | [mx_pool.members.name](data-sources--dns_lb_pool--properties--mx_pool--members.md#schema-mx_pool--members--name) |
| `mx_pool.members.priority` | [mx_pool.members.priority](data-sources--dns_lb_pool--properties--mx_pool--members.md#schema-mx_pool--members--priority) |
| `mx_pool.members.ratio` | [mx_pool.members.ratio](data-sources--dns_lb_pool--properties--mx_pool--members.md#schema-mx_pool--members--ratio) |
| `name` | [name](data-sources--dns_lb_pool--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--dns_lb_pool--reference.md#schema-namespace) |
| `srv_pool` | [srv_pool](data-sources--dns_lb_pool--properties--srv_pool.md#section) |
| `srv_pool.max_answers` | [srv_pool.max_answers](data-sources--dns_lb_pool--properties--srv_pool.md#schema-srv_pool--max_answers) |
| `srv_pool.members` | [srv_pool.members](data-sources--dns_lb_pool--properties--srv_pool--members.md#section) |
| `srv_pool.members.final_translation` | [srv_pool.members.final_translation](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--final_translation) |
| `srv_pool.members.name` | [srv_pool.members.name](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--name) |
| `srv_pool.members.port` | [srv_pool.members.port](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--port) |
| `srv_pool.members.priority` | [srv_pool.members.priority](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--priority) |
| `srv_pool.members.ratio` | [srv_pool.members.ratio](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--ratio) |
| `srv_pool.members.target` | [srv_pool.members.target](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--target) |
| `srv_pool.members.weight` | [srv_pool.members.weight](data-sources--dns_lb_pool--properties--srv_pool--members.md#schema-srv_pool--members--weight) |
| `ttl` | [ttl](data-sources--dns_lb_pool--reference.md#schema-ttl) |
| `use_rrset_ttl` | [use_rrset_ttl](data-sources--dns_lb_pool--properties--use_rrset_ttl.md#section) |

## Next pages

- [a_pool](data-sources--dns_lb_pool--properties--a_pool.md)
- [aaaa_pool](data-sources--dns_lb_pool--properties--aaaa_pool.md)
- [cname_pool](data-sources--dns_lb_pool--properties--cname_pool.md)
- [mx_pool](data-sources--dns_lb_pool--properties--mx_pool.md)
- [srv_pool](data-sources--dns_lb_pool--properties--srv_pool.md)
- [use_rrset_ttl](data-sources--dns_lb_pool--properties--use_rrset_ttl.md)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
