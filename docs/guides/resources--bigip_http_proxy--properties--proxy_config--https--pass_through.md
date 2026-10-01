---
page_title: "proxy_config.https.pass_through"
subcategory: ""
description: "proxy_config.https.pass_through for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1141, "body_sha256": "sha256:63924d123be24f8c2ae3cad9d3d9708212ca1a914b8573c250f975ba4ac31879", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--pass_through.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "pass_through"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/pass_through/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.pass_through for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.pass_through

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- proxy_config.https.pass_through

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
