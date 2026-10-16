---
page_title: "https.tls_cert_options.tls_cert_params.tls_config.default_security"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https tls cert options tls cert params tls config default security"], "body_bytes": 1590, "body_sha256": "sha256:98bd51c45173c9f710ff0e07e7c6c873852107f8c8c5f880638a3792f285330b", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config:default_security", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config", "path": "documentation/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/tls_config/default_security/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0120001133011022-3023333001330231-2022102023023132-0211132313213033-3031200311200322-0030122132223101-3202221322232221-3200120021201321", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_cert_params", "tls_config", "default_security"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/tls_config/default_security/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options.tls_cert_params.tls_config.default_security

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/)
- [https.tls_cert_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/)
- [https.tls_cert_options.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/)
- [https.tls_cert_options.tls_cert_params.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/tls_config/)
- https.tls_cert_options.tls_cert_params.tls_config.default_security

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.
