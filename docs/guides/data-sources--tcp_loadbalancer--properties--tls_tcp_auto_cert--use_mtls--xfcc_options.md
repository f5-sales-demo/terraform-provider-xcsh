---
page_title: "tls_tcp_auto_cert.use_mtls.xfcc_options"
subcategory: "Load Balancing"
description: "tls_tcp_auto_cert.use_mtls.xfcc_options for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2162, "body_sha256": "sha256:9882da4dd01ffe77a93322f2c3e01cbc0888626f01e3c718bd8262c5d7d5d61c", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:xfcc_options", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:xfcc_options", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--xfcc_options.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_tcp_auto_cert", "use_mtls", "xfcc_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_tcp_auto_cert.use_mtls.xfcc_options for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md)
- tls_tcp_auto_cert.use_mtls.xfcc_options

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

<a id="schema-tls_tcp_auto_cert--use_mtls--xfcc_options--xfcc_header_elements"></a>

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

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
