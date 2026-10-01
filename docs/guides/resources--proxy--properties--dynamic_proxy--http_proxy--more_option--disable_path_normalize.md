---
page_title: "dynamic_proxy.http_proxy.more_option.disable_path_normalize"
subcategory: ""
description: "dynamic_proxy.http_proxy.more_option.disable_path_normalize for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1261, "body_sha256": "sha256:7f49d76c049f3b06fd8b0d5a8d4d04c605a94a9c93460de6cb108be1e0149d59", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:disable_path_normalize", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:disable_path_normalize", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--http_proxy--more_option--disable_path_normalize.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "disable_path_normalize"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/disable_path_normalize/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.http_proxy.more_option.disable_path_normalize for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option.disable_path_normalize

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.http_proxy](resources--proxy--properties--dynamic_proxy--http_proxy.md)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--properties--dynamic_proxy--http_proxy--more_option.md)
- dynamic_proxy.http_proxy.more_option.disable_path_normalize

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_path_normalize = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dynamic_proxy.http_proxy.more_option](resources--proxy--properties--dynamic_proxy--http_proxy--more_option.md)
- [xcsh_proxy](../resources/proxy.md)
