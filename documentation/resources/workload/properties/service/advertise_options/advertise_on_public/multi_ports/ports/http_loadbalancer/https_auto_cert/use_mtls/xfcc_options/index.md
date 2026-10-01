---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4512, "body_sha256": "sha256:e12832bc7622987057ca7457fae6a68f78331d89964dad402fe35eeb30b8c3a8", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert:use_mtls", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https_auto_cert/use_mtls/xfcc_options/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "https_auto_cert", "use_mtls", "xfcc_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https_auto_cert/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https_auto_cert/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https_auto_cert/use_mtls/)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
```

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

Terraform syntax:

```terraform
xfcc_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements"></a>

### xfcc_header_elements property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/https_auto_cert/use_mtls/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
