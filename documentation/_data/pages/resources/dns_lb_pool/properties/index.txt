---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": ["dns lb pool"], "body_bytes": 24166, "body_sha256": "sha256:1fe613855b2523e240add42e84b0490150cd03e6a90ef0c42b8d54f6a5bdaea7", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_pool:properties:a_pool", "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool", "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "xcsh-docs:resources:dns_lb_pool:properties:mx_pool", "xcsh-docs:resources:dns_lb_pool:properties:srv_pool", "xcsh-docs:resources:dns_lb_pool:properties:timeouts", "xcsh-docs:resources:dns_lb_pool:properties:use_rrset_ttl"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:reference", "parent_id": "xcsh-docs:resources:dns_lb_pool:fundamentals", "path": "documentation/resources/dns_lb_pool/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["a pool"], "anchor": "section", "description": "Pool for A Record.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "a_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:disable_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "a_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:health_check", "type": "conflicts"}, {"anchor": "schema-a_pool--max_answers", "enforcement": "provider-schema", "group": "a_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "a_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:a_pool:members", "type": "requires"}], "schema_path": ["a_pool"], "syntax": "block", "type": "object"}, {"aliases": ["aaaa pool"], "anchor": "section", "description": "Pool for AAAA Record.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aaaa_pool--max_answers", "enforcement": "provider-schema", "group": "aaaa_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aaaa_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:aaaa_pool:members", "type": "requires"}], "schema_path": ["aaaa_pool"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cname pool"], "anchor": "section", "description": "Pool for CNAME Record.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cname_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cname_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cname_pool:RequiredObjectAttributes:members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "type": "requires"}], "schema_path": ["cname_pool"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["load balancing mode"], "anchor": "schema-load_balancing_mode", "description": "- ROUND_ROBIN: Round-Robin Round Robin will ensure random equal distribution of requests among all pool members in a pool. - RATIO_MEMBER: Ratio-Member Ratio-Member performs load balancing of requests across the pool members based on the ratio assigned to each pool member - STATIC_PERSIST: Static-Persist The Static", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["load_balancing_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["mx pool"], "anchor": "section", "description": "Pool for MX Record.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-mx_pool--max_answers", "enforcement": "provider-schema", "group": "mx_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mx_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:mx_pool:members", "type": "requires"}], "schema_path": ["mx_pool"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["srv pool"], "anchor": "section", "description": "Pool for SRV Record.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-srv_pool--max_answers", "enforcement": "provider-schema", "group": "srv_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "srv_pool:RequiredObjectAttributes:max_answers,members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "type": "requires"}], "schema_path": ["srv_pool"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["ttl"], "anchor": "schema-ttl", "description": "Exclusive with Custom TTL in seconds (default 30) for responses from this pool.", "document_id": "xcsh-docs:resources:dns_lb_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ttl"], "syntax": "attribute", "type": "number"}, {"aliases": ["use rrset ttl"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:use_rrset_ttl", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_rrset_ttl"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_dns_lb_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- Property reference

## Direct properties

- [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/): complete subsection reference.

- [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Type: `"string"`. Optional, Computed.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"),
}
```

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

- [mx_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the DNS LB Pool. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Namespace for the DNS LB Pool. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
}
```

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

- [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/): complete subsection reference.

<a id="schema-ttl"></a>

### ttl property

Type: `"number"`. Optional, Computed.

\[OneOf: ttl, use\_rrset\_ttl\] Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default
&#8203;30) for responses from this pool.

Upstream description:

Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default 30) for responses from this pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

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

- [ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-ttl)
- [use_rrset_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/use_rrset_ttl/#section)

Select alternatives according to the provider validators above.

- [use_rrset_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/use_rrset_ttl/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `a_pool` | [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/#section) |
| `a_pool.disable_health_check` | [a_pool.disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/disable_health_check/#section) |
| `a_pool.health_check` | [a_pool.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/health_check/#section) |
| `a_pool.health_check.name` | [a_pool.health_check.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/health_check/#schema-a_pool--health_check--name) |
| `a_pool.health_check.namespace` | [a_pool.health_check.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/health_check/#schema-a_pool--health_check--namespace) |
| `a_pool.health_check.tenant` | [a_pool.health_check.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/health_check/#schema-a_pool--health_check--tenant) |
| `a_pool.max_answers` | [a_pool.max_answers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/#schema-a_pool--max_answers) |
| `a_pool.members` | [a_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/#section) |
| `a_pool.members.disable_spec` | [a_pool.members.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/#schema-a_pool--members--disable_spec) |
| `a_pool.members.ip_endpoint` | [a_pool.members.ip_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/#schema-a_pool--members--ip_endpoint) |
| `a_pool.members.name` | [a_pool.members.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/#schema-a_pool--members--name) |
| `a_pool.members.priority` | [a_pool.members.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/#schema-a_pool--members--priority) |
| `a_pool.members.ratio` | [a_pool.members.ratio](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/members/#schema-a_pool--members--ratio) |
| `aaaa_pool` | [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/#section) |
| `aaaa_pool.max_answers` | [aaaa_pool.max_answers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/#schema-aaaa_pool--max_answers) |
| `aaaa_pool.members` | [aaaa_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/#section) |
| `aaaa_pool.members.disable_spec` | [aaaa_pool.members.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/#schema-aaaa_pool--members--disable_spec) |
| `aaaa_pool.members.ip_endpoint` | [aaaa_pool.members.ip_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/#schema-aaaa_pool--members--ip_endpoint) |
| `aaaa_pool.members.name` | [aaaa_pool.members.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/#schema-aaaa_pool--members--name) |
| `aaaa_pool.members.priority` | [aaaa_pool.members.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/#schema-aaaa_pool--members--priority) |
| `aaaa_pool.members.ratio` | [aaaa_pool.members.ratio](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/members/#schema-aaaa_pool--members--ratio) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-annotations) |
| `cname_pool` | [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/#section) |
| `cname_pool.disable_health_check` | [cname_pool.disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/disable_health_check/#section) |
| `cname_pool.health_check` | [cname_pool.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/health_check/#section) |
| `cname_pool.health_check.name` | [cname_pool.health_check.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/health_check/#schema-cname_pool--health_check--name) |
| `cname_pool.health_check.namespace` | [cname_pool.health_check.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/health_check/#schema-cname_pool--health_check--namespace) |
| `cname_pool.health_check.tenant` | [cname_pool.health_check.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/health_check/#schema-cname_pool--health_check--tenant) |
| `cname_pool.members` | [cname_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/#section) |
| `cname_pool.members.domain` | [cname_pool.members.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/#schema-cname_pool--members--domain) |
| `cname_pool.members.final_translation` | [cname_pool.members.final_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/#schema-cname_pool--members--final_translation) |
| `cname_pool.members.name` | [cname_pool.members.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/#schema-cname_pool--members--name) |
| `cname_pool.members.priority` | [cname_pool.members.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/#schema-cname_pool--members--priority) |
| `cname_pool.members.ratio` | [cname_pool.members.ratio](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/#schema-cname_pool--members--ratio) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-labels) |
| `load_balancing_mode` | [load_balancing_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-load_balancing_mode) |
| `mx_pool` | [mx_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/#section) |
| `mx_pool.max_answers` | [mx_pool.max_answers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/#schema-mx_pool--max_answers) |
| `mx_pool.members` | [mx_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/#section) |
| `mx_pool.members.domain` | [mx_pool.members.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/#schema-mx_pool--members--domain) |
| `mx_pool.members.name` | [mx_pool.members.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/#schema-mx_pool--members--name) |
| `mx_pool.members.priority` | [mx_pool.members.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/#schema-mx_pool--members--priority) |
| `mx_pool.members.ratio` | [mx_pool.members.ratio](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/members/#schema-mx_pool--members--ratio) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-namespace) |
| `srv_pool` | [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/#section) |
| `srv_pool.max_answers` | [srv_pool.max_answers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/#schema-srv_pool--max_answers) |
| `srv_pool.members` | [srv_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#section) |
| `srv_pool.members.final_translation` | [srv_pool.members.final_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--final_translation) |
| `srv_pool.members.name` | [srv_pool.members.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--name) |
| `srv_pool.members.port` | [srv_pool.members.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--port) |
| `srv_pool.members.priority` | [srv_pool.members.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--priority) |
| `srv_pool.members.ratio` | [srv_pool.members.ratio](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--ratio) |
| `srv_pool.members.target` | [srv_pool.members.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--target) |
| `srv_pool.members.weight` | [srv_pool.members.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/members/#schema-srv_pool--members--weight) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/#schema-timeouts--update) |
| `ttl` | [ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/#schema-ttl) |
| `use_rrset_ttl` | [use_rrset_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/use_rrset_ttl/#section) |

## Next pages

- [a_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/a_pool/)
- [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/aaaa_pool/)
- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/)
- [mx_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/mx_pool/)
- [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/timeouts/)
- [use_rrset_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/use_rrset_ttl/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
