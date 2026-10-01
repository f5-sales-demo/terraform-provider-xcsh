---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3599, "body_sha256": "sha256:c2ef48c4c236bca165048a9add6bf547dc74cf9df00ff66ecc5bf49a66ec4a4f", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:use_mtls:xfcc_options", "child_ids": [], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:use_mtls:xfcc_options", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters:use_mtls", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls--xfcc_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https", "tls_parameters", "use_mtls", "xfcc_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/https/tls_parameters/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="section"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls--xfcc_options--xfcc_header_elements"></a>

### xfcc_header_elements property

Type: `["list", "string"]`. Computed.

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

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters--use_mtls.md)
- [xcsh_workload](../data-sources/workload.md)
