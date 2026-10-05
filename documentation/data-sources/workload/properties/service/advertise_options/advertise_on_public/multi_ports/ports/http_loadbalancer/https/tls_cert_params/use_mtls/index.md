---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls"
subcategory: "Container"
description: "Validation context for downstream client TLS connections."
xcsh_docs: {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls"], "body_bytes": 8480, "body_sha256": "sha256:87cd6c76d127856e59597f7f984f9d1b3eafa7f7e04f25624634517c32113d26", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:crl", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:no_crl", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:trusted_ca", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:xfcc_disabled", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:xfcc_options"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003", "registry_path": "docs/guides/data-sources--workload--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls client certificate optional"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_cert_params--use_mtls--client_certificate_optional", "description": "Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "client_certificate_optional"], "syntax": "attribute", "type": "bool"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls crl"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:crl", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls no crl"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:no_crl", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "no_crl"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls trusted ca"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "trusted_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls trusted ca url"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_cert_params--use_mtls--trusted_ca_url", "description": "Exclusive with Upload a Root CA Certificate specifically for this Load Balancer.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls xfcc disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:xfcc_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "xfcc_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls xfcc options"], "anchor": "section", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls:xfcc_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls", "xfcc_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Validation context for downstream client TLS connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

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

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_cert_params--use_mtls--client_certificate_optional"></a>

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

- [crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/crl/): complete subsection reference.

- [no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/no_crl/): complete subsection reference.

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/trusted_ca/): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_cert_params--use_mtls--trusted_ca_url"></a>

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

- [xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/xfcc_disabled/): complete subsection reference.

- [xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/xfcc_options/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/crl/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/no_crl/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/trusted_ca/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/xfcc_disabled/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/xfcc_options/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
