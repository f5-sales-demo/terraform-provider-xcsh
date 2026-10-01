---
page_title: "dscp_based_queue"
subcategory: ""
description: "dscp_based_queue for xcsh_forwarding_class."
xcsh_docs: {"aliases": [], "body_bytes": 1579, "body_sha256": "sha256:7f1229c9b5b829581f2c9f35ca713a28ca8a6f3e7eaa8853289c2d4559b4fd56", "child_ids": [], "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:resources:forwarding_class:properties:dscp_based_queue", "parent_id": "xcsh-docs:resources:forwarding_class:reference", "path": "documentation/resources/forwarding_class/properties/dscp_based_queue/index.md", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["dscp_based_queue"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/properties/dscp_based_queue/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dscp_based_queue for xcsh_forwarding_class.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dscp_based_queue

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/)
- dscp_based_queue

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

- [dscp_based_queue](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/dscp_based_queue/#section)
- [queue_id_to_use](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/#schema-queue_id_to_use)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dscp_based_queue = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/)
- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/)
