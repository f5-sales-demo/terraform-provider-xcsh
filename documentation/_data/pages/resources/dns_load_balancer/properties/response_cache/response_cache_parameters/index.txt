---
page_title: "response_cache.response_cache_parameters"
subcategory: "DNS"
description: "Configuration parameter for response cache parameters."
xcsh_docs: {"aliases": ["response cache response cache parameters"], "body_bytes": 4195, "body_sha256": "sha256:70ad804de1a1c2d9eb8e2cb343f2360610b29502b3acd5fe122024723fa10457", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "path": "documentation/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3331112003223121-2003203323102330-3202003210020000-1330302321221011-2010113030232121-1220011210121321-1300100132020112-3010132200032130", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [{"anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv6", "enforcement": "provider-schema", "group": "response_cache.response_cache_parameters:RequiredObjectAttributes:cache_cidr_ipv6", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache", "response_cache_parameters"], "schema_version": 1, "sections": [{"aliases": ["response cache response cache parameters cache cidr ipv4"], "anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv4", "description": "Length of CIDR masks used to group IPv4 clients.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_cidr_ipv4"], "syntax": "attribute", "type": "number"}, {"aliases": ["response cache response cache parameters cache cidr ipv6"], "anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv6", "description": "Length of CIDR masks used to group IPv6 clients.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_cidr_ipv6"], "syntax": "attribute", "type": "number"}, {"aliases": ["response cache response cache parameters cache ttl"], "anchor": "schema-response_cache--response_cache_parameters--cache_ttl", "description": "TTL for response cache.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Configuration parameter for response cache parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache.response_cache_parameters

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/)
- response_cache.response_cache_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_cidr_ipv6")}
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
response_cache_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-response_cache--response_cache_parameters--cache_cidr_ipv4"></a>

### cache_cidr_ipv4 property

Type: `"number"`. Optional.

Length of CIDR masks used to group IPv4 clients.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-response_cache--response_cache_parameters--cache_cidr_ipv6"></a>

### cache_cidr_ipv6 property

Type: `"number"`. Optional.

Length of CIDR masks used to group IPv6 clients.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="schema-response_cache--response_cache_parameters--cache_ttl"></a>

### cache_ttl property

Type: `"number"`. Optional.

TTL. TTL for response cache.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```
