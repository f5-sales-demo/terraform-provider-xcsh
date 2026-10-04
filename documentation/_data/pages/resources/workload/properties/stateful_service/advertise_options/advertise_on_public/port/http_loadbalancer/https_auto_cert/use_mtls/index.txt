---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls"
subcategory: "Container"
description: "Validation context for downstream client TLS connections."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls"], "body_bytes": 8320, "body_sha256": "sha256:40ebd5d53b705c5713499dbfcb34398139518c76892fd5c933a36495f27ad0dc", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:crl", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:no_crl", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:trusted_ca", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_disabled", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122", "registry_path": "docs/guides/resources--workload--reference--group-027.md", "relationships": [{"anchor": "schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca_url", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls:ConflictingObjectAttributes:crl,no_crl", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:crl", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls:ConflictingObjectAttributes:crl,no_crl", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:no_crl", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls:ConflictingObjectAttributes:xfcc_disabled,xfcc_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls:ConflictingObjectAttributes:xfcc_disabled,xfcc_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls client certificate optional"], "anchor": "schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--client_certificate_optional", "description": "Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "client_certificate_optional"], "syntax": "attribute", "type": "bool"}, {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls crl"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:crl", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--crl--name", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:crl", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "crl"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls no crl"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:no_crl", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "no_crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls trusted ca"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:trusted_ca", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca--name", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:trusted_ca", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "trusted_ca"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls trusted ca url"], "anchor": "schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca_url", "description": "Exclusive with Upload a Root CA Certificate specifically for this Load Balancer.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls xfcc disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "xfcc_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public port http loadbalancer https auto cert use mtls xfcc options"], "anchor": "section", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options:RequiredObjectAttributes:xfcc_header_elements", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "use_mtls", "xfcc_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Validation context for downstream client TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

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

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--client_certificate_optional"></a>

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

- [crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/crl/): complete subsection reference.

- [no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/no_crl/): complete subsection reference.

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/trusted_ca/): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https_auto_cert--use_mtls--trusted_ca_url"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/xfcc_disabled/): complete subsection reference.

- [xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/xfcc_options/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/crl/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/no_crl/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/trusted_ca/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/xfcc_disabled/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/use_mtls/xfcc_options/)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
