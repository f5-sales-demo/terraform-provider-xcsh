---
page_title: "origin_servers.k8s_service.snat_pool.no_snat_pool"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers k8s service snat pool no snat pool", "upstream servers"], "body_bytes": 1634, "body_sha256": "sha256:0dc714386e3a164a8438f9b8360df00cbd50ecb707fda19256615b5314f83bce", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service:snat_pool", "path": "documentation/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/no_snat_pool/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2021331300301000-0201112021010013-3030120000311113-3213033221203332-1322203033021022-0000202221101311-2200013212221201-1311010002330321", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "k8s_service", "snat_pool", "no_snat_pool"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.k8s_service.snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/)
- [origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/)
- [origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
