---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: "Container"
description: "Header Transformation OPTIONS for HTTP/1.1 request/response headers."
xcsh_docs: {"aliases": ["stateful service advertise options advertise custom ports http loadbalancer https auto cert http protocol options http protocol enable v1 only header transformation"], "body_bytes": 6130, "body_sha256": "sha256:ea796c4f909fc0c70b0ebc6bc6c19d830d65b067741b48966799f6ec3da08271", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212", "registry_path": "docs/guides/data-sources--workload--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "sections": [{"aliases": ["default header transformation"], "anchor": "section", "description": "Use the platform's current default HTTP header transformation behavior.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "default_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["preserve case header transformation"], "anchor": "section", "description": "Preserve HTTP header-name case when upstream case must remain unchanged.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "preserve_case_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["proper case header transformation"], "anchor": "section", "description": "Transform HTTP header names to proper case when explicit transformation is required.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "proper_case_header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [stateful_service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

## Direct properties

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
