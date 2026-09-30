---
page_title: "http_proxy.enable_http"
subcategory: ""
description: "http_proxy.enable_http for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 849, "body_sha256": "sha256:4e19b8cd9e0c21efabe073e4600048aed65121599ba2c734a5b92e19af0bbcd3", "canonical_id": "xcsh-docs:resources:proxy:properties:http_proxy:enable_http", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:enable_http", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy", "path": "docs/guides/resources--proxy--properties--http_proxy--enable_http.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_proxy", "enable_http"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/enable_http/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy.enable_http for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http_proxy.enable_http

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [http_proxy](resources--proxy--properties--http_proxy.md)
- http_proxy.enable_http

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable http.

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
enable_http {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_proxy](resources--proxy--properties--http_proxy.md)
- [xcsh_proxy](../resources/proxy.md)
