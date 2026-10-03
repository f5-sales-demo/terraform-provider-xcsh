---
page_title: "tls_tcp_auto_cert.use_mtls.xfcc_options"
subcategory: "Load Balancing"
description: "X-Forwarded-Client-Cert header elements to be added to requests."
xcsh_docs: {"aliases": ["tls tcp auto cert use mtls xfcc options"], "body_bytes": 2730, "body_sha256": "sha256:04112329b5913b62d874237110c8726fbbd9d206a5178f27cab1cade2913144c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:xfcc_options", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/xfcc_options/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1130133122020122-1203201011332202-0113220130221321-3031122013001103-1211212012303103-2210110030111113-0133120302322110-1110212200100300", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-003.md", "relationships": [{"anchor": "schema-tls_tcp_auto_cert--use_mtls--xfcc_options--xfcc_header_elements", "enforcement": "provider-schema", "group": "tls_tcp_auto_cert.use_mtls.xfcc_options:RequiredObjectAttributes:xfcc_header_elements", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:xfcc_options", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp_auto_cert", "use_mtls", "xfcc_options"], "schema_version": 1, "sections": [{"aliases": ["tls tcp auto cert use mtls xfcc options xfcc header elements"], "anchor": "schema-tls_tcp_auto_cert--use_mtls--xfcc_options--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:use_mtls:xfcc_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp_auto_cert", "use_mtls", "xfcc_options", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "X-Forwarded-Client-Cert header elements to be added to requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [tls_tcp_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/)
- [tls_tcp_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/)
- tls_tcp_auto_cert.use_mtls.xfcc_options

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

<a id="schema-tls_tcp_auto_cert--use_mtls--xfcc_options--xfcc_header_elements"></a>

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

- [tls_tcp_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/use_mtls/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
