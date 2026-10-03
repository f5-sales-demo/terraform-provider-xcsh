---
page_title: "default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation"
subcategory: "Load Balancing"
description: "Transform HTTP header names to proper case when explicit transformation is required."
xcsh_docs: {"aliases": ["default pool advanced options http1 config header transformation proper case header transformation"], "body_bytes": 2103, "body_sha256": "sha256:b26ee0b71f554ebd655d1dd8a638bced4435de05526d846645561fe3f829e985", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation", "path": "documentation/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/proper_case_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1202032102112013-2121032320000201-0232132011133103-3320333120322221-3100322103120321-3310103101031233-3231312132201021-1222133231213032", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "proper_case_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/proper_case_header_transformation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Transform HTTP header names to proper case when explicit transformation is required.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- [default_pool.advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/)
- [default_pool.advanced_options.http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/)
- default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options.http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
