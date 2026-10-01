---
page_title: "tls_tcp_auto_cert"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1508, "body_sha256": "sha256:178615669b06b9f054c3b77263283f4f19df6da9ffd488a906650b8e6a298f9f", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp_auto_cert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- tls_tcp_auto_cert

<a id="section"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with automatic certificates.

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

- [no_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--no_mtls.md): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md): complete subsection reference.

## Next pages

- [tls_tcp_auto_cert.no_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--no_mtls.md)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
