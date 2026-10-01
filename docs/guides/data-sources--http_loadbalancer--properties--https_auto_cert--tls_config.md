---
page_title: "https_auto_cert.tls_config"
subcategory: "Load Balancing"
description: "https_auto_cert.tls_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2227, "body_sha256": "sha256:b1f170d1daace1d9c3d25b2295d40b03a9b44a09fe98257bb178425792ebdebd", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:tls_config", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:tls_config:custom_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:tls_config:default_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:tls_config:low_security", "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:tls_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert", "path": "docs/guides/data-sources--http_loadbalancer--properties--https_auto_cert--tls_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https_auto_cert/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.tls_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https_auto_cert](data-sources--http_loadbalancer--properties--https_auto_cert.md)
- https_auto_cert.tls_config

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

- [custom_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [https_auto_cert.tls_config.custom_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--custom_security.md)
- [https_auto_cert.tls_config.default_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--default_security.md)
- [https_auto_cert.tls_config.low_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--low_security.md)
- [https_auto_cert.tls_config.medium_security](data-sources--http_loadbalancer--properties--https_auto_cert--tls_config--medium_security.md)
- [https_auto_cert](data-sources--http_loadbalancer--properties--https_auto_cert.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
