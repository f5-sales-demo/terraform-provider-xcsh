---
page_title: "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: ""
description: "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3265, "body_sha256": "sha256:0478faa0531bc2d1eafa3092651971f191f9c41251a3d76085ce3b09a6ff27fb", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only.md)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

## Direct properties

- [default_header_transformation](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md): complete subsection reference.

## Next pages

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
