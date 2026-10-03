---
page_title: "tls_parameters.disable_session_key_caching"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters disable session key caching"], "body_bytes": 1239, "body_sha256": "sha256:01ccc08e516f3a8064fd271454045948dfc715d145803bc6515c1df29c8a2213", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_session_key_caching", "parent_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "path": "documentation/data-sources/cluster/properties/tls_parameters/disable_session_key_caching/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3131003313100023-2011121223333200-0211332321102113-3302131121032312-1301111301003330-3020313203010010-1213131322111210-1232103311201020", "registry_path": "docs/guides/data-sources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "disable_session_key_caching"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
