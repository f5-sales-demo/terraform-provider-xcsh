---
page_title: "proxy_config.https.http_protocol_options"
subcategory: ""
description: "proxy_config.https.http_protocol_options for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2268, "body_sha256": "sha256:60b4a714b2bc85bf573932240a74a30cb11bf9f74f41f558c784a2d5b9227dc4", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.http_protocol_options for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.http_protocol_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- proxy_config.https.http_protocol_options

<a id="section"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

## Direct properties

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only.md): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_v2.md): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v2_only.md): complete subsection reference.

## Next pages

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only.md)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_v2.md)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v2_only.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
