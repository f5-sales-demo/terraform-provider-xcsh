---
page_title: "proxy_config.https.tls_cert_params.use_mtls"
subcategory: ""
description: "Validation context for downstream client TLS connections."
xcsh_docs: {"aliases": ["proxy config https tls cert params use mtls"], "body_bytes": 6407, "body_sha256": "sha256:8499558adfc7064a844afd2ab372637c68a3a387b1c31c9636e8e4320f8d43c0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:crl", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:no_crl", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:trusted_ca", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_disabled", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2223212102303322-2012223021222120-2200031110120303-3321102102110220-3233121231230232-0313132312031032-0312232220332222-3010231222013310", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca_url", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls:ConflictingObjectAttributes:crl,no_crl", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:crl", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls:ConflictingObjectAttributes:crl,no_crl", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:no_crl", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls:ConflictingObjectAttributes:xfcc_disabled,xfcc_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls:ConflictingObjectAttributes:xfcc_disabled,xfcc_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["proxy config https tls cert params use mtls client certificate optional"], "anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--client_certificate_optional", "description": "Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "client_certificate_optional"], "syntax": "attribute", "type": "bool"}, {"aliases": ["proxy config https tls cert params use mtls crl"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:crl", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--crl--name", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls.crl:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:crl", "type": "requires"}], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "crl"], "syntax": "block", "type": "object"}, {"aliases": ["proxy config https tls cert params use mtls no crl"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:no_crl", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "no_crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https tls cert params use mtls trusted ca"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:trusted_ca", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--name", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls.trusted_ca:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:trusted_ca", "type": "requires"}], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "trusted_ca"], "syntax": "block", "type": "object"}, {"aliases": ["proxy config https tls cert params use mtls trusted ca url"], "anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca_url", "description": "Exclusive with Upload a Root CA Certificate specifically for this Load Balancer.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy config https tls cert params use mtls xfcc disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "xfcc_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config https tls cert params use mtls xfcc options"], "anchor": "section", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls.xfcc_options:RequiredObjectAttributes:xfcc_header_elements", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options", "type": "requires"}], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "xfcc_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Validation context for downstream client TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params.use_mtls

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/)
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

- [crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/crl/): complete subsection reference.

- [no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/no_crl/): complete subsection reference.

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/trusted_ca/): complete subsection reference.

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

- [xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_disabled/): complete subsection reference.

- [xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_options/): complete subsection reference.

## Next pages

- [proxy_config.https.tls_cert_params.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/crl/)
- [proxy_config.https.tls_cert_params.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/no_crl/)
- [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/trusted_ca/)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_disabled/)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_options/)
- [proxy_config.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
