---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5517, "body_sha256": "sha256:a2d0ee32c77c6d55967fbbca3f252bfc5a0b5322d0f73c4073070421223ff17c", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:certificates", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:no_mtls", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:tls_config", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/certificates/): complete subsection reference.

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/certificates/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/no_mtls/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/tls_config/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
