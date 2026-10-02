---
page_title: "enable_api_discovery.default_api_auth_discovery"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable api discovery default api auth discovery"], "body_bytes": 1359, "body_sha256": "sha256:a60c79b80d821ddd1acf5d2b53be18b6378abedfae4cd55b326a2329ee8ffb41", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "path": "documentation/resources/http_loadbalancer/properties/enable_api_discovery/default_api_auth_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "default_api_auth_discovery"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/default_api_auth_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.default_api_auth_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.default_api_auth_discovery

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
default_api_auth_discovery = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
