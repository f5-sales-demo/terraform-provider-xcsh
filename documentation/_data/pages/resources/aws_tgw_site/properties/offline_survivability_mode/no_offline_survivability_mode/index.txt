---
page_title: "offline_survivability_mode.no_offline_survivability_mode"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["offline survivability mode no offline survivability mode"], "body_bytes": 1125, "body_sha256": "sha256:57efd6cfc8a2ef08bb61c6835c4b3320bcb0a18ad7ad14a27e17b7a5d3e534d4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode:no_offline_survivability_mode", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode", "path": "documentation/resources/aws_tgw_site/properties/offline_survivability_mode/no_offline_survivability_mode/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1211231003331122-3230110231022113-0022001200110322-0310103021323011-0000133310023113-3132000331002203-0332123011312211-1123102122322213", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["offline_survivability_mode", "no_offline_survivability_mode"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/offline_survivability_mode/no_offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode.no_offline_survivability_mode

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/offline_survivability_mode/)
- offline_survivability_mode.no_offline_survivability_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

This is an empty object or choice marker. It has no direct properties.
