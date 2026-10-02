---
page_title: "origin_servers.consul_service.snat_pool.no_snat_pool"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers consul service snat pool no snat pool", "upstream servers"], "body_bytes": 1692, "body_sha256": "sha256:07d5d1c5fc00727bd32a17dfb827c6ee0524f174606342f4f2341b8e81c34f0c", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool", "path": "documentation/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1311301202032301-0010131323121321-3023220101220133-1303221300230203-0331201203223102-0020103130221031-2032003303012010-0220233221133320", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "consul_service", "snat_pool", "no_snat_pool"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.consul_service.snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/)
- [origin_servers.consul_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/)
- origin_servers.consul_service.snat_pool.no_snat_pool

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

- [origin_servers.consul_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
