---
page_title: "auto_http_config"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["auto http config"], "body_bytes": 1335, "body_sha256": "sha256:0c09add6cf6005f8b454fc3a8b07b0cd7f97acc04b1f469b21da585e11d36bce", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:auto_http_config", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/auto_http_config/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0121220033001011-3230302333131331-1120220130301100-2331202102323301-3103232211000023-1201210120201312-0202301131223322-2021003222112331", "registry_path": "docs/guides/data-sources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["auto_http_config"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/auto_http_config/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["clusterCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# auto_http_config

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- auto_http_config

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: auto\_http\_config, http1\_config, http2\_options\] Enable this option

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

OneOf alternatives in this subsection:

- [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/auto_http_config/#section)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/http1_config/#section)
- [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/http2_options/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
