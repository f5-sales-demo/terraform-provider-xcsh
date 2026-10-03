---
page_title: "default_pool.origin_servers.private_name.site_locator"
subcategory: "Load Balancing"
description: "This message defines a reference to a site or virtual site object."
xcsh_docs: {"aliases": ["default pool origin servers private name site locator"], "body_bytes": 2600, "body_sha256": "sha256:3c3f2080fdefb443f64879f133b6398ea0b7507a01e488a302e5de82f19a75ce", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name:site_locator:site", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name:site_locator:virtual_site"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name:site_locator", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/site_locator/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0310320011032122-2311301301130313-0303200313201211-3022331013021321-0233302120202012-1122031222203303-0311100322233033-0301130233121123", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_name", "site_locator"], "schema_version": 1, "sections": [{"aliases": ["default pool origin servers private name site locator site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name:site_locator:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_name", "site_locator", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool origin servers private name site locator virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name:site_locator:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_name", "site_locator", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/site_locator/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This message defines a reference to a site or virtual site object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.private_name.site_locator

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/)
- default_pool.origin_servers.private_name.site_locator

<a id="section"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/site_locator/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/site_locator/virtual_site/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.private_name.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/site_locator/site/)
- [default_pool.origin_servers.private_name.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/site_locator/virtual_site/)
- [default_pool.origin_servers.private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
