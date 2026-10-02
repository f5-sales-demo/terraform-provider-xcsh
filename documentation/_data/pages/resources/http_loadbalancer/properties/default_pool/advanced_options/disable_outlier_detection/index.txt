---
page_title: "default_pool.advanced_options.disable_outlier_detection"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default pool advanced options disable outlier detection"], "body_bytes": 1645, "body_sha256": "sha256:e49dbdcfe00cb42210df475a623421a1ef6dbbb3bc0c7ef471704269fcf0ded8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_outlier_detection", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "path": "documentation/resources/http_loadbalancer/properties/default_pool/advanced_options/disable_outlier_detection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1300111032032202-0223330012333200-0003032323033300-1220012210221033-1221031213013231-0112210331000110-3103001312310113-2332322021231202", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "disable_outlier_detection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/disable_outlier_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.disable_outlier_detection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- default_pool.advanced_options.disable_outlier_detection

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
disable_outlier_detection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
