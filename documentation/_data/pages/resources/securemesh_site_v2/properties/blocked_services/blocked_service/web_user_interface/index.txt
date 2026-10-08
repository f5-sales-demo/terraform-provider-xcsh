---
page_title: "blocked_services.blocked_service.web_user_interface"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["blocked services blocked service web user interface"], "body_bytes": 1236, "body_sha256": "sha256:893c069310c8978b54faca705acea761e3318c5803ca0922728ab8a872b18cdb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:web_user_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service", "path": "documentation/resources/securemesh_site_v2/properties/blocked_services/blocked_service/web_user_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0311113333110132-1332331102333301-1033301322332030-2201030230223030-1032031310230003-2320012320231200-1301031130212110-1031030131310223", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services", "blocked_service", "web_user_interface"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/blocked_services/blocked_service/web_user_interface/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.blocked_service.web_user_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/blocked_services/)
- [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/blocked_services/blocked_service/)
- blocked_services.blocked_service.web_user_interface

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
web_user_interface = {}
```

This is an empty object or choice marker. It has no direct properties.
