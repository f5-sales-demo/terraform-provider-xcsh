---
page_title: "response_cache.response_cache_parameters"
subcategory: "DNS"
description: "Configuration parameter for response cache parameters."
xcsh_docs: {"aliases": ["response cache response cache parameters"], "body_bytes": 3400, "body_sha256": "sha256:c1b344cbf26f2041ce2bc9496e1b8df3c18f42cff77fcc44e79b0447181c5363", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "path": "documentation/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3322233121222221-1232020102211000-0003211203101023-0122200112202212-3100313113023302-2130121120101123-3023221303021123-2101303102300033", "registry_path": "docs/guides/data-sources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache", "response_cache_parameters"], "schema_version": 1, "sections": [{"aliases": ["response cache response cache parameters cache cidr ipv4"], "anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv4", "description": "Length of CIDR masks used to group IPv4 clients.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_cidr_ipv4"], "syntax": "attribute", "type": "number"}, {"aliases": ["response cache response cache parameters cache cidr ipv6"], "anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv6", "description": "Length of CIDR masks used to group IPv6 clients.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_cidr_ipv6"], "syntax": "attribute", "type": "number"}, {"aliases": ["response cache response cache parameters cache ttl"], "anchor": "schema-response_cache--response_cache_parameters--cache_ttl", "description": "TTL for response cache.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:response_cache_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "response_cache_parameters", "cache_ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for response cache parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache.response_cache_parameters

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/)
- response_cache.response_cache_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache parameters.

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

<a id="schema-response_cache--response_cache_parameters--cache_cidr_ipv4"></a>

### cache_cidr_ipv4 property

Type: `"number"`. Computed.

Length of CIDR masks used to group IPv4 clients.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"number"`. Computed.

Length of CIDR masks used to group IPv6 clients.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"number"`. Computed.

TTL. TTL for response cache.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
