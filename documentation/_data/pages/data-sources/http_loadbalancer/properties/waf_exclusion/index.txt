---
page_title: "waf_exclusion"
subcategory: "Load Balancing"
description: "Configuration parameter for waf exclusion."
xcsh_docs: {"aliases": ["waf exclusion"], "body_bytes": 1795, "body_sha256": "sha256:f6caeaf984feb9171cb5190c6ea70a02e3138c207433e602345a93a74bb371f7", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_policy"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/waf_exclusion/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion waf exclusion inline rules"], "anchor": "section", "description": "A list of WAF exclusion rules that will be applied inline.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf exclusion waf exclusion policy"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/waf_exclusion/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for waf exclusion.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- waf_exclusion

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for waf exclusion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

## Direct properties

- [waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/): complete subsection reference.

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_policy/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/)
- [waf_exclusion.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
