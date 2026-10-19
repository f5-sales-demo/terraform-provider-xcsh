---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports"
subcategory: "Container"
description: "Ports to advertise."
xcsh_docs: {"aliases": ["service advertise options advertise on public multi ports ports"], "body_bytes": 2685, "body_sha256": "sha256:7ef5790258371a65138d62ea539a14c03d8ebb6ac14af752ce333fa8f63a90a8", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:port", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:tcp_loadbalancer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122", "registry_path": "docs/guides/resources--workload--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer"], "anchor": "section", "description": "HTTP/HTTPS Load balancer.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports port"], "anchor": "section", "description": "Port of the workload.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "port"], "syntax": "block", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports tcp loadbalancer"], "anchor": "section", "description": "TCP loadbalancer.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:tcp_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "tcp_loadbalancer"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Ports to advertise.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/port/): complete subsection reference.

- [tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/tcp_loadbalancer/): complete subsection reference.
