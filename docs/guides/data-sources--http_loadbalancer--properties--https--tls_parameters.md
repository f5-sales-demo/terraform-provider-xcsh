---
page_title: "https.tls_parameters"
subcategory: "Load Balancing"
description: "https.tls_parameters for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1915, "body_sha256": "sha256:c4a5555c2fba71b9918a6f07be0fe963e4580205c256e32dc397e85e9d25e5cb", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:no_mtls", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_certificates", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:tls_config", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--tls_parameters.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- https.tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

## Direct properties

- [no_mtls](data-sources--http_loadbalancer--properties--https--tls_parameters--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--properties--https--tls_parameters--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls.md): complete subsection reference.

## Next pages

- [https.tls_parameters.no_mtls](data-sources--http_loadbalancer--properties--https--tls_parameters--no_mtls.md)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--properties--https--tls_parameters--tls_certificates.md)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--properties--https--tls_parameters--tls_config.md)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
