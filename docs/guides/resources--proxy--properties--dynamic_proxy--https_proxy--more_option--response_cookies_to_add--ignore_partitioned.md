---
page_title: "dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned"
subcategory: ""
description: "dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1548, "body_sha256": "sha256:3ed483570e43a2a2bb99b287473455e37bfbed8042ee3f9a591473320487cc9e", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add:ignore_partitioned", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add:ignore_partitioned", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:response_cookies_to_add", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_partitioned.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "more_option", "response_cookies_to_add", "ignore_partitioned"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/response_cookies_to_add/ignore_partitioned/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](resources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--properties--dynamic_proxy--https_proxy--more_option.md)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md)
- [xcsh_proxy](../resources/proxy.md)
