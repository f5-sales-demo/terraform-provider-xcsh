---
page_title: "tls_parameters.disable_session_key_caching"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters disable session key caching"], "body_bytes": 988, "body_sha256": "sha256:b94378457a082ee70577020ee484079709ee360e2b89294c559e6c04643ab649", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_session_key_caching", "parent_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "path": "documentation/data-sources/cluster/properties/tls_parameters/disable_session_key_caching/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3131003313100023-2011121223333200-0211332321102113-3302131121032312-1301111301003330-3020313203010010-1213131322111210-1232103311201020", "registry_path": "docs/guides/data-sources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "disable_session_key_caching"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["clusterCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.disable_session_key_caching

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/)
- tls_parameters.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

This is an empty object or choice marker. It has no direct properties.
