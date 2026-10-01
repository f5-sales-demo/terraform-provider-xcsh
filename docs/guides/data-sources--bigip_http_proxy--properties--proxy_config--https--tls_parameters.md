---
page_title: "proxy_config.https.tls_parameters"
subcategory: ""
description: "proxy_config.https.tls_parameters for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2222, "body_sha256": "sha256:d5b2d8eb3a97fee1c940eb87153427f68ade05b254bb8d1f60d57955ee73edea", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:no_mtls", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_parameters for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_parameters

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- proxy_config.https.tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

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

- [no_mtls](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_parameters.no_mtls](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--no_mtls.md)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config.md)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls.md)
- [proxy_config.https](data-sources--bigip_http_proxy--properties--proxy_config--https.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
