---
page_title: "dynamic_proxy.http_proxy"
subcategory: ""
description: "dynamic_proxy.http_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1170, "body_sha256": "sha256:e0cc7f5e9140773383415fb53630224403b7f9f1d977476113c3f5360f2fa973", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--http_proxy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.http_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- dynamic_proxy.http_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Dynamic HTTP Proxy Type. Parameters for dynamic HTTP proxy.

Upstream description:

Parameters for dynamic HTTP proxy.

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
http_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [more_option](resources--proxy--properties--dynamic_proxy--http_proxy--more_option.md): complete subsection reference.

## Next pages

- [dynamic_proxy.http_proxy.more_option](resources--proxy--properties--dynamic_proxy--http_proxy--more_option.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [xcsh_proxy](../resources/proxy.md)
