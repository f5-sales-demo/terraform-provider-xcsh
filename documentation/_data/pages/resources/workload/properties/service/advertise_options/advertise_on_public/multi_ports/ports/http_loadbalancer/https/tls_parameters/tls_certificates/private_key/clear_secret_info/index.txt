---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info"
subcategory: "Container"
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "service advertise options advertise on public multi ports ports http loadbalancer https tls parameters tls certificates private key clear secret info", "tls certificates"], "body_bytes": 6473, "body_sha256": "sha256:73f423e9a6bf428dce718700ac1bd405497a9219e967a2050d9c3b21082f3f5f", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3233131032002012-1023332003300231-2203300311232200-1002033210211020-3013320211001002-3112231303133232-2033211123101332-2320202022010120", "registry_path": "docs/guides/resources--workload--reference--group-010.md", "relationships": [{"anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "private_key", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "private_key", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_parameters:tls_certificates:private_key:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_parameters", "tls_certificates", "private_key", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_certificates/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https--tls_parameters--tls_certificates--private_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_parameters/tls_certificates/private_key/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
