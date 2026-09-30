---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 6108, "body_sha256": "sha256:b0ec5ed847b314c69dfece392752b5798fd511297c575c624f142d672a7fc596", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:http_protocol_options:http_protocol_enable_v1_only", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/http_protocol_options/http_protocol_enable_v1_only/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
