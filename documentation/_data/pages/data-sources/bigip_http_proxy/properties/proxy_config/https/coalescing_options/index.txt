---
page_title: "proxy_config.https.coalescing_options"
subcategory: ""
description: "TLS connection coalescing configuration (not compatible with mTLS)"
xcsh_docs: {"aliases": ["proxy config https coalescing options"], "body_bytes": 2324, "body_sha256": "sha256:bd1e4be7a2bb02710794450540aa6cb6cd9a3e5d9e49e88fb42cd418df31b76b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "coalescing_options"], "schema_version": 1, "sections": [{"aliases": ["default coalescing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "coalescing_options", "default_coalescing"], "syntax": "attribute", "type": "object"}, {"aliases": ["strict coalescing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "coalescing_options", "strict_coalescing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "TLS connection coalescing configuration (not compatible with mTLS)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.coalescing_options

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- proxy_config.https.coalescing_options

<a id="section"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

## Direct properties

- [default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/default_coalescing/): complete subsection reference.

- [strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/strict_coalescing/): complete subsection reference.

## Next pages

- [proxy_config.https.coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/default_coalescing/)
- [proxy_config.https.coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/strict_coalescing/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
