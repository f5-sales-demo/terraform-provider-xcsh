---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options"
subcategory: "Container"
description: "X-Forwarded-Client-Cert header elements to be added to requests."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer https auto cert use mtls xfcc options"], "body_bytes": 4102, "body_sha256": "sha256:4dab78e9cf25212ef68bf4572b7d382ee95a196c2b8536df0b96d52ebd26a988", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:use_mtls", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/use_mtls/xfcc_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0313113021102333-0001033033120201-2221111132100321-0120033003012012-2022223101220022-3330230223013023-3322013300320223-3000211100123010", "registry_path": "docs/guides/resources--workload--reference--group-007.md", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options:RequiredObjectAttributes:xfcc_header_elements", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "use_mtls", "xfcc_options"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise custom ports http loadbalancer https auto cert use mtls xfcc options xfcc header elements"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:https_auto_cert:use_mtls:xfcc_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "https_auto_cert", "use_mtls", "xfcc_options", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "X-Forwarded-Client-Cert header elements to be added to requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/use_mtls/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

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

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements"></a>

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

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/https_auto_cert/use_mtls/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
