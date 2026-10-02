---
page_title: "proxy_config.https.tls_parameters"
subcategory: ""
description: "Inline TLS parameters."
xcsh_docs: {"aliases": ["proxy config https tls parameters"], "body_bytes": 2912, "body_sha256": "sha256:115a249b826bbe528c6269cbf546435c0c85221415144fe52920a337c56d932f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:no_mtls", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:use_mtls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "tls_parameters", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates"], "anchor": "section", "description": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["proxy_config", "https", "tls_parameters", "tls_certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "tls_parameters", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config", "https", "tls_parameters", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Inline TLS parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.tls_parameters

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- proxy_config.https.tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

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

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/no_mtls/): complete subsection reference.

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/): complete subsection reference.

## Next pages

- [proxy_config.https.tls_parameters.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/no_mtls/)
- [proxy_config.https.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/)
- [proxy_config.https.tls_parameters.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/)
- [proxy_config.https.tls_parameters.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
