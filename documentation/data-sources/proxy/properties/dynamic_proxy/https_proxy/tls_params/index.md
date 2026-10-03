---
page_title: "dynamic_proxy.https_proxy.tls_params"
subcategory: ""
description: "Inline TLS parameters."
xcsh_docs: {"aliases": ["dynamic proxy https proxy tls params"], "body_bytes": 2810, "body_sha256": "sha256:89eb2045d31181cba9c1fdcea7f7612d94f195be797a69e5a597ba7c435934d8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133", "registry_path": "docs/guides/data-sources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy tls params no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "dynamic proxy https proxy tls params tls certificates", "existing certificates", "tls certificates"], "anchor": "section", "description": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Inline TLS parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/)
- dynamic_proxy.https_proxy.tls_params

<a id="section"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

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

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/no_mtls/): complete subsection reference.

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.tls_params.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/no_mtls/)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/)
- [dynamic_proxy.https_proxy.tls_params.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_config/)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
