---
page_title: "proxy_config.https.tls_cert_params"
subcategory: ""
description: "Select TLS Parameters and Certificates."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "proxy config https tls cert params", "tls certificates"], "body_bytes": 2928, "body_sha256": "sha256:835ea8e0b4fa525e57ae3c6960bd5d9819570cde7660111c3febeea7d4347b19", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:certificates", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:no_mtls", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "tls_cert_params"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificates", "existing certificates", "tls certificates"], "anchor": "section", "description": "Select one or more certificates with any domain names.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "tls_cert_params", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Select TLS Parameters and Certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_cert_params

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- proxy_config.https.tls_cert_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

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

- [certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/certificates/): complete subsection reference.

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/no_mtls/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/): complete subsection reference.

## Next pages

- [proxy_config.https.tls_cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/certificates/)
- [proxy_config.https.tls_cert_params.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/no_mtls/)
- [proxy_config.https.tls_cert_params.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/)
- [proxy_config.https.tls_cert_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
