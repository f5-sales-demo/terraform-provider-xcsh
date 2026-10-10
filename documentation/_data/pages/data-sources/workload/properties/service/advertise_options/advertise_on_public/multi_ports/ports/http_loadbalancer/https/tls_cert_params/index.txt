---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params"
subcategory: "Container"
description: "Select TLS Parameters and Certificates."
xcsh_docs: {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params"], "body_bytes": 3387, "body_sha256": "sha256:d50524c521922d2824c271882b355089306daf34fbbf02810b06ca90c5ff79e0", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:certificates", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:no_mtls", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:tls_config", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003", "registry_path": "docs/guides/data-sources--workload--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "service advertise options advertise on public multi ports ports http loadbalancer https tls cert params certificates", "tls certificates"], "anchor": "section", "description": "Select one or more certificates with any domain names.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:no_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:tls_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer https tls cert params use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https:tls_cert_params:use_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https", "tls_cert_params", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Select TLS Parameters and Certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

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

## Direct properties

- [certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/certificates/): complete subsection reference.

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https/tls_cert_params/use_mtls/): complete subsection reference.
