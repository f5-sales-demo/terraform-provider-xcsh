---
page_title: "dscp_based_queue"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dscp based queue"], "body_bytes": 1271, "body_sha256": "sha256:82f9561e7150831f770082d4ffb2f90b362569df08b089461d7f0aa73fe23b50", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:properties:dscp_based_queue", "parent_id": "xcsh-docs:data-sources:forwarding_class:reference", "path": "documentation/data-sources/forwarding_class/properties/dscp_based_queue/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3000331101023302-0301122121210323-3001002202033121-2023220023320201-0010023030032003-3030033010033210-1323020033133212-3313113321131301", "registry_path": "docs/guides/data-sources--forwarding_class--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dscp_based_queue"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/properties/dscp_based_queue/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dscp_based_queue

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/)
- dscp_based_queue

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dscp\_based\_queue, queue\_id\_to\_use\] Configuration parameter for dscp based queue.

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

- [dscp_based_queue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp_based_queue/#section)
- [queue_id_to_use](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-queue_id_to_use)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
