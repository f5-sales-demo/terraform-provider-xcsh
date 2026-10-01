---
page_title: "tls_tcp.tls_cert_params"
subcategory: "Load Balancing"
description: "tls_tcp.tls_cert_params for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1952, "body_sha256": "sha256:c97f55813e853dca967661d534187cf8b5ecde1e4dbd96eb6c47bbbe0c1104f1", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:certificates", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp.tls_cert_params for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_cert_params

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [tls_tcp](data-sources--tcp_loadbalancer--properties--tls_tcp.md)
- tls_tcp.tls_cert_params

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

- [certificates](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--certificates.md): complete subsection reference.

- [no_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--no_mtls.md): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md): complete subsection reference.

## Next pages

- [tls_tcp.tls_cert_params.certificates](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--certificates.md)
- [tls_tcp.tls_cert_params.no_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--no_mtls.md)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config.md)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md)
- [tls_tcp](data-sources--tcp_loadbalancer--properties--tls_tcp.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
