---
page_title: "proxy_config.https.tls_cert_params.use_mtls.xfcc_options"
subcategory: ""
description: "X-Forwarded-Client-Cert header elements to be added to requests."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "proxy config https tls cert params use mtls xfcc options", "tls certificates"], "body_bytes": 3149, "body_sha256": "sha256:e80302afa0da738e8bad01680d1dc05c2f031ac8c065db8e4041d39c15301538", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_options/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1120102033101201-0123012302213200-3120300323230003-2111123203122310-0333003320221320-3121320101322321-0123002232323132-0211310201113221", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params.use_mtls.xfcc_options:RequiredObjectAttributes:xfcc_header_elements", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "xfcc_options"], "schema_version": 1, "sections": [{"aliases": ["xfcc header elements"], "anchor": "schema-proxy_config--https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls:xfcc_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls", "xfcc_options", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "X-Forwarded-Client-Cert header elements to be added to requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [proxy_config.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/)
- [proxy_config.https.tls_cert_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

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

<a id="schema-proxy_config--https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements"></a>

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

- [proxy_config.https.tls_cert_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
