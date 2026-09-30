---
page_title: "proxy_config.https.tls_cert_params.use_mtls"
subcategory: ""
description: "proxy_config.https.tls_cert_params.use_mtls for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 5484, "body_sha256": "sha256:acf2b1ac26c41db1c0dd2e10cf5213e015938d8c08ac951f140f3658ce3ff0c2", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:crl", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:no_crl", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:trusted_ca", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_disabled", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https.tls_cert_params.use_mtls for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https.tls_cert_params.use_mtls

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- proxy_config.https.tls_cert_params.use_mtls

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

<a id="schema-proxy_config--https--tls_cert_params--use_mtls--client_certificate_optional"></a>

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

- [crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--crl.md): complete subsection reference.

- [no_crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--no_crl.md): complete subsection reference.

- [trusted_ca](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--trusted_ca.md): complete subsection reference.

<a id="schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca_url"></a>

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

- [xfcc_disabled](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_disabled.md): complete subsection reference.

- [xfcc_options](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_options.md): complete subsection reference.

## Next pages

- [proxy_config.https.tls_cert_params.use_mtls.crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--crl.md)
- [proxy_config.https.tls_cert_params.use_mtls.no_crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--no_crl.md)
- [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--trusted_ca.md)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_disabled.md)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_options.md)
- [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
