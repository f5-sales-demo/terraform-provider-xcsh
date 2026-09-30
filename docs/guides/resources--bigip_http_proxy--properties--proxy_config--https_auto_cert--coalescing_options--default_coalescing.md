---
page_title: "proxy_config.https_auto_cert.coalescing_options.default_coalescing"
subcategory: ""
description: "proxy_config.https_auto_cert.coalescing_options.default_coalescing for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1350, "body_sha256": "sha256:d1d2852935932375a6619b6dc6e35715f99e8a3e673240d7df55dc0fc5e5cb69", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options:default_coalescing", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options:default_coalescing", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:coalescing_options", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--default_coalescing.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "coalescing_options", "default_coalescing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/coalescing_options/default_coalescing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https_auto_cert.coalescing_options.default_coalescing for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https_auto_cert.coalescing_options.default_coalescing

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https_auto_cert](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md)
- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options.md)
- proxy_config.https_auto_cert.coalescing_options.default_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
