---
page_title: "https.tls_cert_options.tls_cert_params.tls_config"
subcategory: "Load Balancing"
description: "https.tls_cert_options.tls_cert_params.tls_config for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2710, "body_sha256": "sha256:6df0261669050aa4fbf64001440f7fafb885ffda2bf5d660eb420cfd9bae1c23", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config:custom_security", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config:default_security", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config:low_security", "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:tls_config", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_cert_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_options.tls_cert_params.tls_config for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https.tls_cert_options.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [https](data-sources--cdn_loadbalancer--properties--https.md)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md)
- https.tls_cert_options.tls_cert_params.tls_config

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

- [custom_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [https.tls_cert_options.tls_cert_params.tls_config.custom_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--custom_security.md)
- [https.tls_cert_options.tls_cert_params.tls_config.default_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--default_security.md)
- [https.tls_cert_options.tls_cert_params.tls_config.low_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--low_security.md)
- [https.tls_cert_options.tls_cert_params.tls_config.medium_security](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--tls_config--medium_security.md)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
