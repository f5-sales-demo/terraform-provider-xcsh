---
page_title: "https.tls_cert_options.tls_cert_params.no_mtls"
subcategory: "Load Balancing"
description: "https.tls_cert_options.tls_cert_params.no_mtls for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1307, "body_sha256": "sha256:81ad98deb28be26806fcaa0268b76267dffe6396e10d081811f281a40814446d", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:no_mtls", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params:no_mtls", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:https:tls_cert_options:tls_cert_params", "path": "docs/guides/resources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params--no_mtls.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_cert_params", "no_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/https/tls_cert_options/tls_cert_params/no_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_options.tls_cert_params.no_mtls for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options.tls_cert_params.no_mtls

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [https](resources--cdn_loadbalancer--properties--https.md)
- [https.tls_cert_options](resources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md)
- https.tls_cert_options.tls_cert_params.no_mtls

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
no_mtls = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--properties--https--tls_cert_options--tls_cert_params.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
