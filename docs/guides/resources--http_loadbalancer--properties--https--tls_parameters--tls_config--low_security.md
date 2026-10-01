---
page_title: "https.tls_parameters.tls_config.low_security"
subcategory: "Load Balancing"
description: "https.tls_parameters.tls_config.low_security for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:1a7fb9d1c3d09a6e0cf99cc58956dc634ae5f3273b1f2d990bb8d40da26e9e85", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config", "path": "docs/guides/resources--http_loadbalancer--properties--https--tls_parameters--tls_config--low_security.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters", "tls_config", "low_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters.tls_config.low_security for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_config.low_security

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [https.tls_parameters](resources--http_loadbalancer--properties--https--tls_parameters.md)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--properties--https--tls_parameters--tls_config.md)
- https.tls_parameters.tls_config.low_security

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
low_security = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https.tls_parameters.tls_config](resources--http_loadbalancer--properties--https--tls_parameters--tls_config.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
