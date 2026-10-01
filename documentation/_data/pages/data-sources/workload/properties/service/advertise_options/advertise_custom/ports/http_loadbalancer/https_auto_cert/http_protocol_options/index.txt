---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4450, "body_sha256": "sha256:b9898aa961d8d0ca4f7dd332a877e61f2855ac86abe723531c98110aa9faa9e4", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v1_v2", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options:http_protocol_enable_v2_only"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:http_protocol_options", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "http_protocol_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v1_v2/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/http_protocol_options/http_protocol_enable_v2_only/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
