---
page_title: "dynamic_proxy.https_proxy"
subcategory: ""
description: "dynamic_proxy.https_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:81ca5d7882fa53cf3c0f38f9faca5d946dc6367508c246a8990a88e65ec481bb", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:more_option", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy--https_proxy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- dynamic_proxy.https_proxy

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for https proxy.

Upstream description:

Parameters for dynamic HTTPS proxy.

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

## Direct properties

- [more_option](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md): complete subsection reference.

- [tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [xcsh_proxy](../data-sources/proxy.md)
