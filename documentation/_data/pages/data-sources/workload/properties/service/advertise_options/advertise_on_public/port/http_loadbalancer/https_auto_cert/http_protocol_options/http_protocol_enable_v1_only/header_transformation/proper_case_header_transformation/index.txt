---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation"
subcategory: "Container"
description: "Transform HTTP header names to proper case when explicit transformation is required."
xcsh_docs: {"aliases": ["service advertise options advertise on public port http loadbalancer https auto cert http protocol options http protocol enable v1 only header transformation proper case header transformation"], "body_bytes": 3878, "body_sha256": "sha256:3b80b913e356abcaf3b4740b4f2328ab69226e133a630290b49470572cdb7e94", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1221033122110031-0001110021331223-0300232231131202-0332303220100131-3311201313330230-0331201211022211-1020120021001022-3110300201233213", "registry_path": "docs/guides/data-sources--workload--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "proper_case_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Transform HTTP header names to proper case when explicit transformation is required.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
