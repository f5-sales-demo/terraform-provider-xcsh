---
page_title: "https.coalescing_options"
subcategory: "Load Balancing"
description: "TLS connection coalescing configuration (not compatible with mTLS)"
xcsh_docs: {"aliases": ["https coalescing options"], "body_bytes": 2050, "body_sha256": "sha256:28d1b79c600c354a941686be4ad1b73494400996634b24c4140bc8d44d87a76d", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https:coalescing_options:default_coalescing", "xcsh-docs:data-sources:http_loadbalancer:properties:https:coalescing_options:strict_coalescing"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:coalescing_options", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https", "path": "documentation/data-sources/http_loadbalancer/properties/https/coalescing_options/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "coalescing_options"], "schema_version": 1, "sections": [{"aliases": ["default coalescing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:coalescing_options:default_coalescing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "coalescing_options", "default_coalescing"], "syntax": "attribute", "type": "object"}, {"aliases": ["strict coalescing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:coalescing_options:strict_coalescing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "coalescing_options", "strict_coalescing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "TLS connection coalescing configuration (not compatible with mTLS)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.coalescing_options

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/)
- https.coalescing_options

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

- [default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/coalescing_options/default_coalescing/): complete subsection reference.

- [strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/coalescing_options/strict_coalescing/): complete subsection reference.

## Next pages

- [https.coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/coalescing_options/default_coalescing/)
- [https.coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/coalescing_options/strict_coalescing/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
