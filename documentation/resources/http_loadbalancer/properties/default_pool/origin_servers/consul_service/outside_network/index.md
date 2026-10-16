---
page_title: "default_pool.origin_servers.consul_service.outside_network"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default pool origin servers consul service outside network"], "body_bytes": 1437, "body_sha256": "sha256:256e00d5f5ab085ffc33400c21c72c6cb6c46bc6a86211c1ae5b5adee77e3015", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service:outside_network", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/outside_network/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-2211012211311332-0231231002020013-1011033211320212-2320300030231020-1130220012033121-3130000031030213-1031212001111112-1233202130213220", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "consul_service", "outside_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/outside_network/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.consul_service.outside_network

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/)
- default_pool.origin_servers.consul_service.outside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

Additional upstream details:

This can be used for messages where no values are needed.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.
