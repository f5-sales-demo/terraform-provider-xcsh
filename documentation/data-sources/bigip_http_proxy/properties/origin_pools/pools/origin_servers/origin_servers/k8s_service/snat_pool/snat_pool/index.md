---
page_title: "origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool"
subcategory: ""
description: "List of IPv4 prefixes that represent an endpoint."
xcsh_docs: {"aliases": ["origin pools pools origin servers origin servers k8s service snat pool snat pool"], "body_bytes": 3054, "body_sha256": "sha256:2cb2f04e27d84face868055ac9da0f5e7c7d204fa69cfb2891b7771a01ff7ca3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool", "path": "documentation/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1313213220013222-0223033121012202-3103220100120033-1013232031031330-3210122032310102-1102300333101312-0201002020030202-3001031321112020", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "k8s_service", "snat_pool", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["origin pools pools origin servers origin servers k8s service snat pool snat pool prefixes"], "anchor": "schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "k8s_service", "snat_pool", "snat_pool", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "List of IPv4 prefixes that represent an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
