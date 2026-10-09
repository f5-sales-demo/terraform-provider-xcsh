---
page_title: "cloudfront.disable_js_insert"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cloudfront disable js insert"], "body_bytes": 1049, "body_sha256": "sha256:22fef9680e843bacc84bbb48ee5e7bc2ec0a22eb32754ccaab365d3516ef8513", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:disable_js_insert", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/disable_js_insert/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2021213232222233-1330301300310122-1020102001112303-2300100211113202-0323221313013012-1033210213010223-1122330233301300-1222030203220202", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "disable_js_insert"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/disable_js_insert/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.disable_js_insert

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.disable_js_insert

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.
