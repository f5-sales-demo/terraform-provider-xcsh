---
page_title: "dscp_based_queue"
subcategory: ""
description: "dscp_based_queue for xcsh_forwarding_class."
xcsh_docs: {"aliases": [], "body_bytes": 1226, "body_sha256": "sha256:d6583b35f293845ad71371c1496d2b59ee7f64d6d1512d45f8f89014a449a607", "canonical_id": "xcsh-docs:data-sources:forwarding_class:properties:dscp_based_queue", "child_ids": [], "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:properties:dscp_based_queue", "parent_id": "xcsh-docs:data-sources:forwarding_class:reference", "path": "docs/guides/data-sources--forwarding_class--properties--dscp_based_queue.md", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dscp_based_queue"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/properties/dscp_based_queue/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dscp_based_queue for xcsh_forwarding_class.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dscp_based_queue

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md)
- [Property reference](data-sources--forwarding_class--reference.md)
- dscp_based_queue

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dscp\_based\_queue, queue\_id\_to\_use\] Configuration parameter for dscp based queue.

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

OneOf alternatives in this subsection:

- [dscp_based_queue](data-sources--forwarding_class--properties--dscp_based_queue.md#section)
- [queue_id_to_use](data-sources--forwarding_class--reference.md#schema-queue_id_to_use)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--forwarding_class--reference.md)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md)
