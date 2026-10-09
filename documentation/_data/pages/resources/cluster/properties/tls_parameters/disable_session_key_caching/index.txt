---
page_title: "tls_parameters.disable_session_key_caching"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters disable session key caching"], "body_bytes": 1049, "body_sha256": "sha256:514bb3a541f53029c448b9071837b68a45b043fbc6fb3c6cfe3a609a2a9767dd", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "path": "documentation/resources/cluster/properties/tls_parameters/disable_session_key_caching/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2102323013232123-2122130023210330-2223011133213113-3010120223212310-3033220231110101-2003103300202020-3311031203312302-0210012033113310", "registry_path": "docs/guides/resources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "disable_session_key_caching"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["clusterCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.disable_session_key_caching

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/)
- tls_parameters.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.
