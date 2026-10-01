---
page_title: "tls_tcp_auto_cert.tls_config"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert.tls_config for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2897, "body_sha256": "sha256:2664e4946e642feae6cfa326179bdf9d4ac53c12d8405e88f2214accfea8af45", "child_ids": ["xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:custom_security", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:default_security", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:low_security", "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/index.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["tls_tcp_auto_cert", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert.tls_config for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.tls_config

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [tls_tcp_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/)
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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/medium_security/): complete subsection reference.

## Next pages

- [tls_tcp_auto_cert.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/custom_security/)
- [tls_tcp_auto_cert.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/default_security/)
- [tls_tcp_auto_cert.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/low_security/)
- [tls_tcp_auto_cert.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/medium_security/)
- [tls_tcp_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
