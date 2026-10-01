---
page_title: "tls_tcp.tls_parameters"
subcategory: "Load Balancing"
description: "tls_tcp.tls_parameters for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1936, "body_sha256": "sha256:77cf99acc72e7c62583da6cbdb43d9560939879759f4aeb37cafa3eadbd0acb1", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:no_mtls", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_config", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp.tls_parameters for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [tls_tcp](data-sources--tcp_loadbalancer--properties--tls_tcp.md)
- tls_tcp.tls_parameters

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

- [no_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [tls_tcp.tls_parameters.no_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--no_mtls.md)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config.md)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md)
- [tls_tcp](data-sources--tcp_loadbalancer--properties--tls_tcp.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
