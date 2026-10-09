---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options"
subcategory: "Container"
description: "HTTP protocol configuration OPTIONS for downstream connections."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer https auto cert http protocol options"], "body_bytes": 2990, "body_sha256": "sha256:d6e42cc6caf8eacbb89dfbfef905aa8026bdd4b20cc37fd2b59b063f30aec139", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133", "registry_path": "docs/guides/data-sources--workload--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise custom ports http loadbalancer https auto cert http protocol options http protocol enable v1 only"], "anchor": "section", "description": "HTTP/1.1 Protocol OPTIONS for downstream connections.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https auto cert http protocol options http protocol enable v1 v2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_v2"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise custom ports http loadbalancer https auto cert http protocol options http protocol enable v2 only"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v2_only"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "HTTP protocol configuration OPTIONS for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="section"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

## Direct properties

- [http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/): complete subsection reference.

- [http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_v2/): complete subsection reference.

- [http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v2_only/): complete subsection reference.
