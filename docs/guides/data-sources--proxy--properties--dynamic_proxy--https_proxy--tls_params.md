---
page_title: "dynamic_proxy.https_proxy.tls_params"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2021, "body_sha256": "sha256:19102904127cf5e88f297052a6a50265273b239a2d0e0251c4048bbce934faeb", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dynamic_proxy.https_proxy.tls_params

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md)
- dynamic_proxy.https_proxy.tls_params

<a id="section"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

## Direct properties

- [no_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.no_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--no_mtls.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md)
- [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md)
- [xcsh_proxy](../data-sources/proxy.md)
