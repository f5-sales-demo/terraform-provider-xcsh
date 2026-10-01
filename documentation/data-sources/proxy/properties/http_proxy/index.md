---
page_title: "http_proxy"
subcategory: ""
description: "http_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1583, "body_sha256": "sha256:255b87998be3aa0e7e5a4cd4121e22b3fd4ecf02d2891d3b8738a373da2aefc8", "child_ids": ["xcsh-docs:data-sources:proxy:properties:http_proxy:enable_http", "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:http_proxy", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "documentation/data-sources/proxy/properties/http_proxy/index.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["http_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/http_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- http_proxy

<a id="section"></a>

Type: `"single"`. Computed.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Upstream description:

Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

## Direct properties

- [enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/enable_http/): complete subsection reference.

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/): complete subsection reference.

## Next pages

- [http_proxy.enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/enable_http/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/http_proxy/more_option/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
