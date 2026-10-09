---
page_title: "dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dynamic proxy https proxy more option response cookies to add ignore domain"], "body_bytes": 1596, "body_sha256": "sha256:d04b4f847de883fe949da57c7f6a2b259cdf5155baf0286174bc70eba65f7b6e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add:ignore_domain", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/ignore_domain/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2130212002130011-3311123202110101-1333211303332300-2131100332110020-1022320212311020-3331310330033132-3111021202100330-3032323232312311", "registry_path": "docs/guides/resources--proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "more_option", "response_cookies_to_add", "ignore_domain"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/ignore_domain/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

This is an empty object or choice marker. It has no direct properties.
