---
page_title: "https.tls_cert_options"
subcategory: "Load Balancing"
description: "https.tls_cert_options for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1497, "body_sha256": "sha256:dff509cdcd71a8245ee65869aea59558b6d5254ce46e13d4f05c91cf2f92ee14", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_options for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [https](data-sources--cdn_loadbalancer--properties--https.md)
- https.tls_cert_options

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert options.

Upstream description:

TLS Certificate OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

## Direct properties

- [tls_cert_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md): complete subsection reference.

- [tls_inline_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_inline_params.md): complete subsection reference.

## Next pages

- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_inline_params.md)
- [https](data-sources--cdn_loadbalancer--properties--https.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
