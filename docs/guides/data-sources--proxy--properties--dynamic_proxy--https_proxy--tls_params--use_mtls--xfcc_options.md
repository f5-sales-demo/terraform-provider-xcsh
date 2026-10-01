---
page_title: "dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2416, "body_sha256": "sha256:bcaf7643ad5646400d692a350a369ffbc3acfa29126919944f59cfdcadd6b293", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_options", "child_ids": [], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_options", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--xfcc_options.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "xfcc_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options

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

<a id="schema-dynamic_proxy--https_proxy--tls_params--use_mtls--xfcc_options--xfcc_header_elements"></a>

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

- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md)
- [xcsh_proxy](../data-sources/proxy.md)
