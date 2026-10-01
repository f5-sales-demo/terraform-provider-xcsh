---
page_title: "https.tls_parameters.tls_config"
subcategory: "Load Balancing"
description: "https.tls_parameters.tls_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3082, "body_sha256": "sha256:963a33acefb75c1bfc24f47f6d8ce756e3f88d08d47ee125c618bb278b635ec9", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters", "path": "documentation/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["https", "tls_parameters", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters.tls_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/)
- [https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/)
- https.tls_parameters.tls_config

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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/medium_security/): complete subsection reference.

## Next pages

- [https.tls_parameters.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/custom_security/)
- [https.tls_parameters.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/default_security/)
- [https.tls_parameters.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/)
- [https.tls_parameters.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/tls_config/medium_security/)
- [https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
