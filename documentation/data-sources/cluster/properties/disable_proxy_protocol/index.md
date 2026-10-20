---
page_title: "disable_proxy_protocol"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable proxy protocol"], "body_bytes": 1462, "body_sha256": "sha256:eff359abecddc0d720d601cc787dfe430b32f3ee1bf6ff672bd848a16d003d03", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:disable_proxy_protocol", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/disable_proxy_protocol/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3032021323222023-2113331010130133-1220333211120201-1312220223131131-1322033303023321-2223300130100302-1001220312222332-1013112312332131", "registry_path": "docs/guides/data-sources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_proxy_protocol"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/disable_proxy_protocol/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["clusterCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_proxy_protocol

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- disable_proxy_protocol

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/disable_proxy_protocol/#section)
- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/proxy_protocol_v1/#section)
- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/proxy_protocol_v2/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
