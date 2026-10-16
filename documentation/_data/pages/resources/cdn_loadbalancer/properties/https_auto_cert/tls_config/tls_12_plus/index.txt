---
page_title: "https_auto_cert.tls_config.tls_12_plus"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https auto cert tls config tls 12 plus"], "body_bytes": 1201, "body_sha256": "sha256:e933511a7617e45b436bba86dc01f5eb133a8b7bb4bd724e259b9b56bfc55b28", "capabilities": ["cdn", "load-balancing.tls"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config:tls_12_plus", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https_auto_cert:tls_config", "path": "documentation/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_12_plus/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1103322121320103-2102011000203232-0233313222010210-0003211002221230-2320303121320331-1213013030231213-1232303021202221-0022320100112132", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert", "tls_config", "tls_12_plus"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/tls_12_plus/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.tls_config.tls_12_plus

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/)
- [https_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/https_auto_cert/tls_config/)
- https_auto_cert.tls_config.tls_12_plus

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls 12 plus.

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
tls_12_plus = {}
```

This is an empty object or choice marker. It has no direct properties.
