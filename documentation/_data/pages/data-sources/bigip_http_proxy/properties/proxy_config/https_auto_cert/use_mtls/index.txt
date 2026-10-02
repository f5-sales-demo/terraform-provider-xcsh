---
page_title: "proxy_config.https_auto_cert.use_mtls"
subcategory: ""
description: "Validation context for downstream client TLS connections."
xcsh_docs: {"aliases": ["proxy config https auto cert use mtls"], "body_bytes": 5599, "body_sha256": "sha256:0e0e679d4ee28df9096ead34180a9bd59cb5bbe8b49bc739bdc4b19a9e0c609c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:crl", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:no_crl", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:trusted_ca", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:xfcc_disabled", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:xfcc_options"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1112313302100303-2032300331001320-2003221231112012-1211210332212020-0212133012203301-0102220003302101-3022121110222030-3000130033323200", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "client certificate optional", "existing certificates", "tls certificates"], "anchor": "schema-proxy_config--https_auto_cert--use_mtls--client_certificate_optional", "description": "Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "client_certificate_optional"], "syntax": "attribute", "type": "bool"}, {"aliases": ["crl"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:crl", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["no crl"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:no_crl", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "no_crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["trusted ca"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "trusted_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca url"], "anchor": "schema-proxy_config--https_auto_cert--use_mtls--trusted_ca_url", "description": "Exclusive with Upload a Root CA Certificate specifically for this Load Balancer.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["xfcc disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:xfcc_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "xfcc_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["xfcc options"], "anchor": "section", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:xfcc_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "xfcc_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Validation context for downstream client TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https_auto_cert.use_mtls

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- proxy_config.https_auto_cert.use_mtls

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

<a id="schema-proxy_config--https_auto_cert--use_mtls--client_certificate_optional"></a>

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

- [crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/crl/): complete subsection reference.

- [no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/no_crl/): complete subsection reference.

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/trusted_ca/): complete subsection reference.

<a id="schema-proxy_config--https_auto_cert--use_mtls--trusted_ca_url"></a>

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

- [xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_disabled/): complete subsection reference.

- [xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_options/): complete subsection reference.

## Next pages

- [proxy_config.https_auto_cert.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/crl/)
- [proxy_config.https_auto_cert.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/no_crl/)
- [proxy_config.https_auto_cert.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/trusted_ca/)
- [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_disabled/)
- [proxy_config.https_auto_cert.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_options/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
