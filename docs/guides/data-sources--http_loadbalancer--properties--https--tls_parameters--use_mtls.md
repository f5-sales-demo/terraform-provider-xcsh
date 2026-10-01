---
page_title: "https.tls_parameters.use_mtls"
subcategory: "Load Balancing"
description: "https.tls_parameters.use_mtls for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4633, "body_sha256": "sha256:a65272ba316dac82c82e8d5130f1e03b44cc8f4b4c0bfcf31346a766ed82af25", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:crl", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:no_crl", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:trusted_ca", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:xfcc_disabled", "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:xfcc_options"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_parameters", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/tls_parameters/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_parameters.use_mtls for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.use_mtls

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- [https.tls_parameters](data-sources--http_loadbalancer--properties--https--tls_parameters.md)
- https.tls_parameters.use_mtls

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

<a id="schema-https--tls_parameters--use_mtls--client_certificate_optional"></a>

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

- [crl](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--crl.md): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--no_crl.md): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--trusted_ca.md): complete subsection reference.

<a id="schema-https--tls_parameters--use_mtls--trusted_ca_url"></a>

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

- [xfcc_disabled](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--xfcc_disabled.md): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--xfcc_options.md): complete subsection reference.

## Next pages

- [https.tls_parameters.use_mtls.crl](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--crl.md)
- [https.tls_parameters.use_mtls.no_crl](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--no_crl.md)
- [https.tls_parameters.use_mtls.trusted_ca](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--trusted_ca.md)
- [https.tls_parameters.use_mtls.xfcc_disabled](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--xfcc_disabled.md)
- [https.tls_parameters.use_mtls.xfcc_options](data-sources--http_loadbalancer--properties--https--tls_parameters--use_mtls--xfcc_options.md)
- [https.tls_parameters](data-sources--http_loadbalancer--properties--https--tls_parameters.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
