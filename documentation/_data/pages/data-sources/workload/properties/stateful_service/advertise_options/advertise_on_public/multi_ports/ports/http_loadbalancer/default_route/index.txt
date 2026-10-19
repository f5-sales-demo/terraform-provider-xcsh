---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route"
subcategory: "Container"
description: "Default route matching all APIs."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer default route"], "body_bytes": 3745, "body_sha256": "sha256:1006d019fd2cc4e70a98585facfdd3ac28b2bef0c2e14e0337cf7f9ba62b610e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route:auto_host_rewrite", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route:disable_host_rewrite"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/default_route/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322", "registry_path": "docs/guides/data-sources--workload--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "default_route"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer default route auto host rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route:auto_host_rewrite", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "default_route", "auto_host_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer default route disable host rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route:disable_host_rewrite", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "default_route", "disable_host_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer default route host rewrite"], "anchor": "schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--default_route--host_rewrite", "description": "Exclusive with Host header will be swapped with this value.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "default_route", "host_rewrite"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/default_route/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Default route matching all APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Additional upstream details:

Default route matching all APIs.

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

## Direct properties

- [auto_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/default_route/auto_host_rewrite/): complete subsection reference.

- [disable_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/default_route/disable_host_rewrite/): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--default_route--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
