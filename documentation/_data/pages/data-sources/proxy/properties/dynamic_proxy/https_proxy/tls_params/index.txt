---
page_title: "dynamic_proxy.https_proxy.tls_params"
subcategory: ""
description: "Inline TLS parameters."
xcsh_docs: {"aliases": ["dynamic proxy https proxy tls params"], "body_bytes": 1757, "body_sha256": "sha256:bbdaa81611a08943f2534ac40fb9c8616e222a8f8cf9f68df4970e96b06ac978", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133", "registry_path": "docs/guides/data-sources--proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy tls params no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:no_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "dynamic proxy https proxy tls params tls certificates", "existing certificates", "tls certificates"], "anchor": "section", "description": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Inline TLS parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
