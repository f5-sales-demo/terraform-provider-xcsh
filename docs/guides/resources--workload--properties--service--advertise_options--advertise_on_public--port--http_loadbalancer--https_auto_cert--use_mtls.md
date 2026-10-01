---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 7091, "body_sha256": "sha256:276143c097db6e9ba7a07013864781300552ff8b45c00f1131f8f45ae7c39de8", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:crl", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:no_crl", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:trusted_ca", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_disabled", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](resources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--client_certificate_optional"></a>

### client_certificate_optional property

Type: `"bool"`. Optional.

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

- [crl](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--crl.md): complete subsection reference.

- [no_crl](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--no_crl.md): complete subsection reference.

- [trusted_ca](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--xfcc_disabled.md): complete subsection reference.

- [xfcc_options](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--crl.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--no_crl.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--xfcc_disabled.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert.md)
- [xcsh_workload](../resources/workload.md)
