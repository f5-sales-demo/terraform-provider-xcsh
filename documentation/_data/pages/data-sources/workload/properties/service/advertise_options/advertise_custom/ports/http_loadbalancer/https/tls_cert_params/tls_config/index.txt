---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5157, "body_sha256": "sha256:d0ed943a9d48b312a39d386248fe2d8066bc12583c411430e5a25a6730fddd81", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:custom_security", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:default_security", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:low_security", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params:tls_config", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https:tls_cert_params", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https", "tls_cert_params", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/medium_security/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/custom_security/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/default_security/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/low_security/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/tls_config/medium_security/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https/tls_cert_params/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
