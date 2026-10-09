---
page_title: "purge_all"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["purge all"], "body_bytes": 836, "body_sha256": "sha256:817d2f08baaae4af11ef80ad8e2e63b232fc6a71371000e890fb64dec691dd99", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_purge_command:properties:purge_all", "parent_id": "xcsh-docs:resources:cdn_purge_command:reference", "path": "documentation/resources/cdn_purge_command/properties/purge_all/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1123320230111100-2132200320220211-2020231302000131-1312020213123222-0321132332223200-2232110213232300-0230310022012131-3223130101212223", "registry_path": "docs/guides/resources--cdn_purge_command--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["purge_all"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/properties/purge_all/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# purge_all

Breadcrumbs:

- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/properties/)
- purge_all

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
purge_all = {}
```

This is an empty object or choice marker. It has no direct properties.
