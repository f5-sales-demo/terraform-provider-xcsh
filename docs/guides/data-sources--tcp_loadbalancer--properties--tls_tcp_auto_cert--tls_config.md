---
page_title: "tls_tcp_auto_cert.tls_config"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert.tls_config for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2248, "body_sha256": "sha256:5b1b96d412d6448d07b4ac2362150f23931450a3bc976b91fccfa7de53387f1b", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:custom_security", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:default_security", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:low_security", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp_auto_cert", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert.tls_config for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.tls_config

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md)
- tls_tcp_auto_cert.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [tls_tcp_auto_cert.tls_config.custom_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--custom_security.md)
- [tls_tcp_auto_cert.tls_config.default_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--default_security.md)
- [tls_tcp_auto_cert.tls_config.low_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--low_security.md)
- [tls_tcp_auto_cert.tls_config.medium_security](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--medium_security.md)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
