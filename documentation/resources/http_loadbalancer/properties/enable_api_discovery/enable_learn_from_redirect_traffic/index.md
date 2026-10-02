---
page_title: "enable_api_discovery.enable_learn_from_redirect_traffic"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable api discovery enable learn from redirect traffic"], "body_bytes": 1428, "body_sha256": "sha256:0b89bcf7a0e01eb0fb7e1135948d4a2aabd0aa1b09c389cffc13edd0e895d379", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "path": "documentation/resources/http_loadbalancer/properties/enable_api_discovery/enable_learn_from_redirect_traffic/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "enable_learn_from_redirect_traffic"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/enable_learn_from_redirect_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.enable_learn_from_redirect_traffic

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
