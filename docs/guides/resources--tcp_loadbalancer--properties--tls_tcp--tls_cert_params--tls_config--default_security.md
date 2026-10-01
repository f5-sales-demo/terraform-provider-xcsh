---
page_title: "tls_tcp.tls_cert_params.tls_config.default_security"
subcategory: "Load Balancing"
description: "tls_tcp.tls_cert_params.tls_config.default_security for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1316, "body_sha256": "sha256:61bc54cef1f4692e1f07359cc718b787a31c5d63283343ed97235e493c63fff1", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:tls_config:default_security", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:tls_config:default_security", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_cert_params:tls_config", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--default_security.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp", "tls_cert_params", "tls_config", "default_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_cert_params/tls_config/default_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp.tls_cert_params.tls_config.default_security for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_cert_params.tls_config.default_security

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params.md)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config.md)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
default_security = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
