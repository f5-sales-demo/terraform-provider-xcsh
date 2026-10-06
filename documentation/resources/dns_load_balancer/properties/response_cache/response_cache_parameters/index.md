---
page_title: "response_cache.response_cache_parameters"
subcategory: "DNS"
description: "Configuration parameter for response cache parameters."
xcsh_docs: {"aliases": ["response cache response cache parameters"], "body_bytes": 4195, "body_sha256": "sha256:525cc0155778635e53087484787a85b24302607f7cf2de55c6d89e3b9349c2bb", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "path": "documentation/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3331112003223121-2003203323102330-3202003210020000-1330302321221011-2010113030232121-1220011210121321-1300100132020112-3010132200032130", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [{"anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv6", "enforcement": "provider-schema", "group": "response_cache.response_cache_parameters:RequiredObjectAttributes:cache_cidr_ipv6", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache", "response_cache_parameters"], "schema_version": 1, "sections": [{"aliases": ["response cache response cache parameters cache cidr ipv4"], "anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv4", "description": "Length of CIDR masks used to group IPv4 clients.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_cidr_ipv4"], "syntax": "attribute", "type": "number"}, {"aliases": ["response cache response cache parameters cache cidr ipv6"], "anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv6", "description": "Length of CIDR masks used to group IPv6 clients.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_cidr_ipv6"], "syntax": "attribute", "type": "number"}, {"aliases": ["response cache response cache parameters cache ttl"], "anchor": "schema-response_cache--response_cache_parameters--cache_ttl", "description": "TTL for response cache.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configuration parameter for response cache parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
