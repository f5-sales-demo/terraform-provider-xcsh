---
page_title: "more_option.response_cookies_to_add.ignore_value"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["more option response cookies to add ignore value"], "body_bytes": 1238, "body_sha256": "sha256:525a29e79cd6aabd878c3a8737d222ab323511299fd6d287ca4189acd2ab4904", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add:ignore_value", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:more_option:response_cookies_to_add", "path": "documentation/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_value/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3100111032103113-2230310102122122-0313103010101213-0223103301010022-0102110020033202-2223130032033032-3300020311131322-1201010312202110", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option", "response_cookies_to_add", "ignore_value"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/ignore_value/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.response_cookies_to_add.ignore_value

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/)
- [more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/more_option/response_cookies_to_add/)
- more_option.response_cookies_to_add.ignore_value

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
