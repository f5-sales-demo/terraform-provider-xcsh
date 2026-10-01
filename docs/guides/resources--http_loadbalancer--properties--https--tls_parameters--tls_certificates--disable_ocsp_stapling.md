---
page_title: "https.tls_parameters.tls_certificates.disable_ocsp_stapling"
subcategory: "Load Balancing"
description: "https.tls_parameters.tls_certificates.disable_ocsp_stapling for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1380, "body_sha256": "sha256:6efa48ba65f229c9de5059022f77bd742cb325d10bc8e4e9bae9e514d35fef86", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates:disable_ocsp_stapling", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_certificates", "path": "docs/guides/resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates--disable_ocsp_stapling.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters.tls_certificates.disable_ocsp_stapling for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [https.tls_parameters](resources--http_loadbalancer--properties--https--tls_parameters.md)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
