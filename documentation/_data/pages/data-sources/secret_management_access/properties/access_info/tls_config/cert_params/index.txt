---
page_title: "access_info.tls_config.cert_params"
subcategory: ""
description: "Certificate Parameters for authentication, TLS ciphers, and trust store."
xcsh_docs: {"aliases": ["access info tls config cert params", "authentication", "cert", "certificate", "credential setup", "credentials", "existing certificates", "tls certificates"], "body_bytes": 8744, "body_sha256": "sha256:b3a9266560d86968f63992c82338550ae073faac74d1a0bb808f566b95ed881d", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:skip_server_verification", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config", "path": "documentation/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params"], "schema_version": 1, "sections": [{"aliases": ["authentication", "cert", "certificate", "certificates", "credential setup", "credentials", "existing certificates", "tls certificates"], "anchor": "section", "description": "Client TLS Certificate required for mTLS authentication.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["cipher suites"], "anchor": "schema-access_info--tls_config--cert_params--cipher_suites", "description": "The following list specifies the supported cipher suite TLS_AES_128_GCM_SHA256 TLS_AES_256_GCM_SHA384 TLS_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "cipher_suites"], "syntax": "attribute", "type": "list"}, {"aliases": ["maximum protocol version"], "anchor": "schema-access_info--tls_config--cert_params--maximum_protocol_version", "description": "TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "maximum_protocol_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["minimum protocol version"], "anchor": "schema-access_info--tls_config--cert_params--minimum_protocol_version", "description": "TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "minimum_protocol_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["skip server verification"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:skip_server_verification", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "skip_server_verification"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls validation params"], "anchor": "section", "description": "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:volterra_trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Certificate Parameters for authentication, TLS ciphers, and trust store.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- access_info.tls_config.cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

## Direct properties

- [certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/): complete subsection reference.

<a id="schema-access_info--tls_config--cert_params--cipher_suites"></a>

### cipher_suites property

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-access_info--tls_config--cert_params--maximum_protocol_version"></a>

### maximum_protocol_version property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--tls_config--cert_params--minimum_protocol_version"></a>

### minimum_protocol_version property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/skip_server_verification/): complete subsection reference.

- [tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/volterra_trusted_ca/): complete subsection reference.

## Next pages

- [access_info.tls_config.cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/)
- [access_info.tls_config.cert_params.skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/skip_server_verification/)
- [access_info.tls_config.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/)
- [access_info.tls_config.cert_params.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/volterra_trusted_ca/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
