---
page_title: "default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "default pool origin servers k8s service snat pool no snat pool", "origin servers", "upstream servers"], "body_bytes": 1979, "body_sha256": "sha256:10dd7f3b6dcf59ba58b02e6048a214b0aef7e91b5ba05be7991969e1fc532498", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/no_snat_pool/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1111320223302200-0330212323101202-1220220210100210-1223130010122030-0200001313233000-2033303122203313-0111212110200310-0121133212221311", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool", "no_snat_pool"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/)
- [default_pool.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/)
- default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

Upstream description:

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
no_snat_pool = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
