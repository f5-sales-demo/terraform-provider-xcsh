---
page_title: "https.tls_cert_params"
subcategory: "Load Balancing"
description: "https.tls_cert_params for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2236, "body_sha256": "sha256:6e982f887d7468cf4834464a71e4156400f3ed1f825f8aa5678548b91c65c6ad", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_cert_params", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:tls_cert_params:certificates", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_cert_params:no_mtls", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_cert_params:tls_config", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_cert_params", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https", "path": "docs/guides/resources--http_loadbalancer--properties--https--tls_cert_params.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_params for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_params

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- https.tls_cert_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
```

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

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [certificates](resources--http_loadbalancer--properties--https--tls_cert_params--certificates.md): complete subsection reference.

- [no_mtls](resources--http_loadbalancer--properties--https--tls_cert_params--no_mtls.md): complete subsection reference.

- [tls_config](resources--http_loadbalancer--properties--https--tls_cert_params--tls_config.md): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--properties--https--tls_cert_params--use_mtls.md): complete subsection reference.

## Next pages

- [https.tls_cert_params.certificates](resources--http_loadbalancer--properties--https--tls_cert_params--certificates.md)
- [https.tls_cert_params.no_mtls](resources--http_loadbalancer--properties--https--tls_cert_params--no_mtls.md)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--properties--https--tls_cert_params--tls_config.md)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--properties--https--tls_cert_params--use_mtls.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
