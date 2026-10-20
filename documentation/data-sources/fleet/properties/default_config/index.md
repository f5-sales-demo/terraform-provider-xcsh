---
page_title: "default_config"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default config"], "body_bytes": 1338, "body_sha256": "sha256:c6eb00164d20f03ffe6fe500c7dd1ac7f9f225ef4225788138c348c0c0887485", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:default_config", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/default_config/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1232232122223230-0300310232023033-2100332200122130-0211200101321332-0000311133222122-0230121123110011-0201012012212022-1221302112223212", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_config"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/default_config/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_config

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- default_config

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_config, device\_list, interface\_list; Default: default\_config\] Enable this
option

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

- [default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/default_config/#section)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/#section)
- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/interface_list/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
