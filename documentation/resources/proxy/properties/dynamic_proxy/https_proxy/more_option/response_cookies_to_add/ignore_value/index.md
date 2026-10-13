---
page_title: "dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dynamic proxy https proxy more option response cookies to add ignore value"], "body_bytes": 1592, "body_sha256": "sha256:a267e63d6044e5fc47dbf073f6a94adc5054716bac8bdc689a56a3ba1f0ed368", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add:ignore_value", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/ignore_value/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "resources", "registry_anchor": "canonical-1313012130201032-2222121001201222-2300123313203311-0221001020222032-0033020033131021-0231111013123123-2100123102122101-0233003112002200", "registry_path": "docs/guides/resources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "more_option", "response_cookies_to_add", "ignore_value"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/ignore_value/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["proxyCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

This is an empty object or choice marker. It has no direct properties.
