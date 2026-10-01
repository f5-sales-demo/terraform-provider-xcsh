---
page_title: "http_proxy.more_option.no_request_limit_per_connection"
subcategory: ""
description: "http_proxy.more_option.no_request_limit_per_connection for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1149, "body_sha256": "sha256:5b523295b1af7c6c9c1ba1b4d66d9fbd9e77f7e93fc4d9238c4a37f7477f8319", "canonical_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:no_request_limit_per_connection", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option", "path": "docs/guides/resources--proxy--properties--http_proxy--more_option--no_request_limit_per_connection.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_proxy", "more_option", "no_request_limit_per_connection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy.more_option.no_request_limit_per_connection for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [http_proxy](resources--proxy--properties--http_proxy.md)
- [http_proxy.more_option](resources--proxy--properties--http_proxy--more_option.md)
- http_proxy.more_option.no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

Upstream description:

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
no_request_limit_per_connection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_proxy.more_option](resources--proxy--properties--http_proxy--more_option.md)
- [xcsh_proxy](../resources/proxy.md)
