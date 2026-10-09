---
page_title: "offline_survivability_mode.enable_offline_survivability_mode"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["offline survivability mode enable offline survivability mode"], "body_bytes": 1165, "body_sha256": "sha256:8877f9abe8431b1445c25b2b2b3e7931d9cd3db7dce5c1207579f5f242066ceb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:offline_survivability_mode:enable_offline_survivability_mode", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:offline_survivability_mode", "path": "documentation/resources/securemesh_site_v2/properties/offline_survivability_mode/enable_offline_survivability_mode/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2203323101230022-1133103332103223-2322333231131013-3322223200313333-0320203103020033-3313002110322230-0310000220213131-1211112300030132", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["offline_survivability_mode", "enable_offline_survivability_mode"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/offline_survivability_mode/enable_offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode.enable_offline_survivability_mode

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/offline_survivability_mode/)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

This is an empty object or choice marker. It has no direct properties.
