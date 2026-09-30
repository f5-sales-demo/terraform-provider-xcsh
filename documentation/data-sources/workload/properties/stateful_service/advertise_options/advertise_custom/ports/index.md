---
page_title: "stateful_service.advertise_options.advertise_custom.ports"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3376, "body_sha256": "sha256:34e81607dbcadb97e413c7c3e1855668d78dbf5c203f61c9f54658057aca16c5", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:tcp_loadbalancer"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_custom.ports

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- stateful_service.advertise_options.advertise_custom.ports

<a id="section"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Direct properties

- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/): complete subsection reference.

- [tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/tcp_loadbalancer/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_custom.ports.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/)
- [stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/tcp_loadbalancer/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
