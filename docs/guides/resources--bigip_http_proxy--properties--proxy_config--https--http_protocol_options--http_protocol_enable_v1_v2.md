---
page_title: "proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2"
subcategory: ""
description: "proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1320, "body_sha256": "sha256:c5e0a3424b0a506356414aebf7fa2907727bf215cf02983744997e56e766bb6f", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_v2.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "http_protocol_options", "http_protocol_enable_v1_v2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_v2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
