---
page_title: "trusted_clients.waf_skip_processing"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["trusted clients waf skip processing"], "body_bytes": 1308, "body_sha256": "sha256:e6a9fecf4a852f76279153a48f8c9bee6e4c084417429d5b38650358da7169b8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients:waf_skip_processing", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:trusted_clients", "path": "documentation/resources/http_loadbalancer/properties/trusted_clients/waf_skip_processing/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0230200100110103-0330023321103022-2112322010330320-0212120122121323-3130132223211233-0010201332230200-3102203100220300-1231303111321002", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["trusted_clients", "waf_skip_processing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/trusted_clients/waf_skip_processing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# trusted_clients.waf_skip_processing

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/trusted_clients/)
- trusted_clients.waf_skip_processing

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
waf_skip_processing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/trusted_clients/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
