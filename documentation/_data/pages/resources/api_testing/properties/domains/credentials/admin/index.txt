---
page_title: "domains.credentials.admin"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["domains credentials admin"], "body_bytes": 1092, "body_sha256": "sha256:27ad2c15806cfe9903aea137a2c83a8abf3f859679e92c5b933c0033e70c72ce", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "path": "documentation/resources/api_testing/properties/domains/credentials/admin/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3122213202111132-2220200131230222-0133133211302110-1003333230010112-1301022313033321-3313323120100023-0323021032320223-2301233300101212", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "admin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/admin/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["api_testingCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.admin

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- domains.credentials.admin

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
admin = {}
```

This is an empty object or choice marker. It has no direct properties.
