---
page_title: "tls_tcp_auto_cert.tls_config.default_security"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert.tls_config.default_security for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1200, "body_sha256": "sha256:d1285d0bb9ab28ff22b5f1fc0096e5638a05602d10431c044f11d653c9d0818e", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:default_security", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:default_security", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "path": "docs/guides/resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--default_security.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp_auto_cert", "tls_config", "default_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/default_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert.tls_config.default_security for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.tls_config.default_security

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md)
- tls_tcp_auto_cert.tls_config.default_security

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

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
