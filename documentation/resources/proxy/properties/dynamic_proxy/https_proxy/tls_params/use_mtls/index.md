---
page_title: "dynamic_proxy.https_proxy.tls_params.use_mtls"
subcategory: ""
description: "Validation context for downstream client TLS connections."
xcsh_docs: {"aliases": ["dynamic proxy https proxy tls params use mtls"], "body_bytes": 4203, "body_sha256": "sha256:e6cbd3bd163057447c7b7dba9b85f86a7c7cbf13830145591b9f5871524f4225", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:crl", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:no_crl", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:trusted_ca", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_disabled", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032", "registry_path": "docs/guides/resources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy tls params use mtls client certificate optional"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--use_mtls--client_certificate_optional", "description": "Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "client_certificate_optional"], "syntax": "attribute", "type": "bool"}, {"aliases": ["dynamic proxy https proxy tls params use mtls crl"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:crl", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "crl"], "syntax": "block", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params use mtls no crl"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:no_crl", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "no_crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params use mtls trusted ca"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca"], "syntax": "block", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params use mtls trusted ca url"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca_url", "description": "Exclusive with Upload a Root CA Certificate specifically for this Load Balancer.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["dynamic proxy https proxy tls params use mtls xfcc disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "xfcc_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params use mtls xfcc options"], "anchor": "section", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_options", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "xfcc_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Validation context for downstream client TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.use_mtls

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/)
- dynamic_proxy.https_proxy.tls_params.use_mtls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dynamic_proxy--https_proxy--tls_params--use_mtls--client_certificate_optional"></a>

### client_certificate_optional property

Type: `"bool"`. Optional.

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

- [crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/crl/): complete subsection reference.

- [no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/no_crl/): complete subsection reference.

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/trusted_ca/): complete subsection reference.

<a id="schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/xfcc_disabled/): complete subsection reference.

- [xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/xfcc_options/): complete subsection reference.
