---
page_title: "https.tls_cert_params.use_mtls.xfcc_options"
subcategory: "Load Balancing"
description: "https.tls_cert_params.use_mtls.xfcc_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2275, "body_sha256": "sha256:1a58778a7fa5e0a09e49047311b09503c7083ddcd6b4f095b3994b08ecbad4af", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:use_mtls:xfcc_options", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:use_mtls:xfcc_options", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_cert_params:use_mtls", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--tls_cert_params--use_mtls--xfcc_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_params", "use_mtls", "xfcc_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/tls_cert_params/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_params.use_mtls.xfcc_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_params.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- [https.tls_cert_params](data-sources--http_loadbalancer--properties--https--tls_cert_params.md)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--properties--https--tls_cert_params--use_mtls.md)
- https.tls_cert_params.use_mtls.xfcc_options

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

<a id="schema-https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements"></a>

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

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--properties--https--tls_cert_params--use_mtls.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
