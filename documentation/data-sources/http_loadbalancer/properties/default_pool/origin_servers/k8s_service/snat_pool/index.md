---
page_title: "default_pool.origin_servers.k8s_service.snat_pool"
subcategory: "Load Balancing"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["default pool origin servers k8s service snat pool"], "body_bytes": 2531, "body_sha256": "sha256:e9ad576a9e133975cfc4df0b91a2505d87ec886a9160029df242cc42f1ae0b40", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1213311303022021-0001331121110013-2012122022233320-2012330323110121-2233123112021010-1222231033101223-1332322131233003-0001102201212133", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["default pool origin servers k8s service snat pool no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool origin servers k8s service snat pool snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool", "snat_pool"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service.snat_pool

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/)
- default_pool.origin_servers.k8s_service.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

## Direct properties

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/snat_pool/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/no_snat_pool/)
- [default_pool.origin_servers.k8s_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/snat_pool/)
- [default_pool.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
