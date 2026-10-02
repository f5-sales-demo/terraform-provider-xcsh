---
page_title: "single_lb_app.disable_discovery"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["single lb app disable discovery"], "body_bytes": 1318, "body_sha256": "sha256:346b0fb23b28acd0a596674fdfc1fb631d5343d8fe5fa311927b017e64f2e246", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/disable_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1002330301111021-3230331100121210-1223300130323020-0213302300120331-0112313302210030-1131331102223311-3223320113133133-2312310122000303", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "disable_discovery"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/disable_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.disable_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- single_lb_app.disable_discovery

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable discovery.

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
disable_discovery = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
