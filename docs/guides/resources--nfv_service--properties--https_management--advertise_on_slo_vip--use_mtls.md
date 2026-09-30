---
page_title: "https_management.advertise_on_slo_vip.use_mtls"
subcategory: ""
description: "https_management.advertise_on_slo_vip.use_mtls for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 5369, "body_sha256": "sha256:3a68d253e7b7f532922529f350683a856b86c9ff3f9bf53fd62d3f4ba6770dd3", "canonical_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls", "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:crl", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:no_crl", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:trusted_ca", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:xfcc_disabled", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:xfcc_options"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip", "path": "docs/guides/resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_vip", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_vip/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_vip.use_mtls for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_management.advertise_on_slo_vip.use_mtls

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_slo_vip](resources--nfv_service--properties--https_management--advertise_on_slo_vip.md)
- https_management.advertise_on_slo_vip.use_mtls

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

<a id="schema-https_management--advertise_on_slo_vip--use_mtls--client_certificate_optional"></a>

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

- [crl](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--crl.md): complete subsection reference.

- [no_crl](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--no_crl.md): complete subsection reference.

- [trusted_ca](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--trusted_ca.md): complete subsection reference.

<a id="schema-https_management--advertise_on_slo_vip--use_mtls--trusted_ca_url"></a>

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

- [xfcc_disabled](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_disabled.md): complete subsection reference.

- [xfcc_options](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_options.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_vip.use_mtls.crl](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--crl.md)
- [https_management.advertise_on_slo_vip.use_mtls.no_crl](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--no_crl.md)
- [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--trusted_ca.md)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_disabled.md)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_options.md)
- [https_management.advertise_on_slo_vip](resources--nfv_service--properties--https_management--advertise_on_slo_vip.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
