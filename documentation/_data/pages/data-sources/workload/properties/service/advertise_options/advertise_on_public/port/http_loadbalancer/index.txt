---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer"
subcategory: "Container"
description: "HTTP/HTTPS Load balancer."
xcsh_docs: {"aliases": ["service advertise options advertise on public port http loadbalancer"], "body_bytes": 7078, "body_sha256": "sha256:11adc1bc6a20b09a1a8f30dc6b4d51c55a382cd01d9332ba0005323897047f2d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:http", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200", "registry_path": "docs/guides/data-sources--workload--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public port http loadbalancer default route"], "anchor": "section", "description": "Default route matching all APIs.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "default_route"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public port http loadbalancer domains"], "anchor": "schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--domains", "description": "A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are supported in the suffix or prefix form Domain search order: 1. Exact domain names: ``www.example.com``. 2. Prefix domain wildcards: ``*.example.com`` or ``*.bar.example.com``. 3. Special wildcard ``*`` matching any", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["service advertise options advertise on public port http loadbalancer http"], "anchor": "section", "description": "Choice for selecting HTTP proxy.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:http", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "http"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public port http loadbalancer https"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https"], "syntax": "attribute", "type": "object"}, {"aliases": ["automatic certificate management", "automatic certificates", "managed TLS certificates", "service advertise options advertise on public port http loadbalancer https auto cert"], "anchor": "section", "description": "Choice for selecting HTTP proxy with bring your own certificates.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public port http loadbalancer specific routes"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HTTP/HTTPS Load balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/)
- service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

## Direct properties

- [default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/http/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/): complete subsection reference.

- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/): complete subsection reference.

- [specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/http/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/)
- [service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
