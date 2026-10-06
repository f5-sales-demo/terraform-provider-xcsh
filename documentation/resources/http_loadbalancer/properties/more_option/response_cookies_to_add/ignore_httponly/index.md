---
page_title: "more_option.response_cookies_to_add.ignore_httponly"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["more option response cookies to add ignore httponly"], "body_bytes": 1250, "body_sha256": "sha256:8c4d489645c1e95ae693fec45ff6fe0562f2910e70eaba9008643cf0a5fc90bc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_httponly", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add", "path": "documentation/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_httponly/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3320123211213101-1112001020132233-0322212101320131-0030131211333122-0320223001223122-2312230112030311-0311303013321322-2110221023233003", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option", "response_cookies_to_add", "ignore_httponly"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_httponly/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.response_cookies_to_add.ignore_httponly

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/)
- [more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/)
- more_option.response_cookies_to_add.ignore_httponly

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.
