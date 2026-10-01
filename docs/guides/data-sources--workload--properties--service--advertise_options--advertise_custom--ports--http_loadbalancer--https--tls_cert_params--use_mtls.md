---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 6860, "body_sha256": "sha256:b47b03a0917616105d6446a2a3ec27e032d405f6f38357f76a2b6c8be18d5497", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls:crl", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls:no_crl", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls:trusted_ca", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls:xfcc_disabled", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls:xfcc_options"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:use_mtls", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](data-sources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="section"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

## Direct properties

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--client_certificate_optional"></a>

### client_certificate_optional property

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--crl.md): complete subsection reference.

- [no_crl](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--no_crl.md): complete subsection reference.

- [trusted_ca](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--trusted_ca.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--xfcc_disabled.md): complete subsection reference.

- [xfcc_options](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--xfcc_options.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--crl.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--no_crl.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--trusted_ca.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--xfcc_disabled.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params--use_mtls--xfcc_options.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--https--tls_cert_params.md)
- [xcsh_workload](../data-sources/workload.md)
