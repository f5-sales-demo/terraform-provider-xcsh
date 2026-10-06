---
page_title: "dynamic_proxy.https_proxy"
subcategory: ""
description: "Parameters for dynamic HTTPS proxy."
xcsh_docs: {"aliases": ["dynamic proxy https proxy"], "body_bytes": 1198, "body_sha256": "sha256:316645881ebc6471b3c2edfeec197947616b9de9ad24b9021e80219e6bda32fb", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:more_option", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "path": "documentation/data-sources/proxy/properties/dynamic_proxy/https_proxy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220", "registry_path": "docs/guides/data-sources--proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy more option"], "anchor": "section", "description": "This defines various OPTIONS to define a route.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "more_option"], "syntax": "attribute", "type": "object"}, {"aliases": ["dynamic proxy https proxy tls params"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:https_proxy:tls_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/https_proxy/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Parameters for dynamic HTTPS proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/)
- dynamic_proxy.https_proxy

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for https proxy.

Additional upstream details:

Parameters for dynamic HTTPS proxy.

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

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/more_option/): complete subsection reference.

- [tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/dynamic_proxy/https_proxy/tls_params/): complete subsection reference.
