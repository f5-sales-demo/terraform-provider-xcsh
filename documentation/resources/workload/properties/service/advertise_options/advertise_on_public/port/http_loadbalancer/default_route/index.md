---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route"
subcategory: "Container"
description: "Default route matching all APIs."
xcsh_docs: {"aliases": ["service advertise options advertise on public port http loadbalancer default route"], "body_bytes": 4912, "body_sha256": "sha256:80db57e061933b672d5f75683616e79a7b4f82ce6a5dd8d9e207332044291217", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:auto_host_rewrite", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:disable_host_rewrite"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013", "registry_path": "docs/guides/resources--workload--reference--group-012.md", "relationships": [{"anchor": "schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--default_route--host_rewrite", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--default_route--host_rewrite", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route:ConflictingObjectAttributes:disable_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route:ConflictingObjectAttributes:auto_host_rewrite,disable_host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:auto_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:auto_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route:ConflictingObjectAttributes:auto_host_rewrite,disable_host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:disable_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route:ConflictingObjectAttributes:disable_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:disable_host_rewrite", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "default_route"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public port http loadbalancer default route auto host rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:auto_host_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "default_route", "auto_host_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public port http loadbalancer default route disable host rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route:disable_host_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "default_route", "disable_host_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public port http loadbalancer default route host rewrite"], "anchor": "schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--default_route--host_rewrite", "description": "Exclusive with Host header will be swapped with this value.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:default_route", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "default_route", "host_rewrite"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Default route matching all APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/auto_host_rewrite/): complete subsection reference.

- [disable_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/disable_host_rewrite/): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--default_route--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/auto_host_rewrite/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/default_route/disable_host_rewrite/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
