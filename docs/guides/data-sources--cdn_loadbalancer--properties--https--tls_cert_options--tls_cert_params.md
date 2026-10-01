---
page_title: "https.tls_cert_options.tls_cert_params"
subcategory: "Load Balancing"
description: "https.tls_cert_options.tls_cert_params for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2296, "body_sha256": "sha256:eabd3c7d4a05652f8d2af935a5e90b8fe80a20fa3b9bf1e48a79a2a32f71d5fa", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:certificates", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:no_mtls", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_options.tls_cert_params for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options.tls_cert_params

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [https](data-sources--cdn_loadbalancer--properties--https.md)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- https.tls_cert_options.tls_cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

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

- [certificates](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--certificates.md): complete subsection reference.

- [no_mtls](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--no_mtls.md): complete subsection reference.

- [tls_config](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--use_mtls.md): complete subsection reference.

## Next pages

- [https.tls_cert_options.tls_cert_params.certificates](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--certificates.md)
- [https.tls_cert_options.tls_cert_params.no_mtls](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--no_mtls.md)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config.md)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--use_mtls.md)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
