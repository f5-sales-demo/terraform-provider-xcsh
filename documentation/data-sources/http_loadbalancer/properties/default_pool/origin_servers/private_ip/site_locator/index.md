---
page_title: "default_pool.origin_servers.private_ip.site_locator"
subcategory: "Load Balancing"
description: "This message defines a reference to a site or virtual site object."
xcsh_docs: {"aliases": ["backend servers", "default pool origin servers private ip site locator", "origin servers", "upstream servers"], "body_bytes": 2576, "body_sha256": "sha256:48d8849e14220b08da765fe52ec2cd6ee5010a72a070e36ae74605a8fb2a6465", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:site", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0120002023322211-3032122001223103-3031322302112111-0310022223012103-3223232323333212-0233011302022002-3011020312333322-2023013000023111", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_ip", "site_locator"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "site_locator", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "site_locator", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This message defines a reference to a site or virtual site object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.private_ip.site_locator

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/)
- default_pool.origin_servers.private_ip.site_locator

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/virtual_site/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.private_ip.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/site/)
- [default_pool.origin_servers.private_ip.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/virtual_site/)
- [default_pool.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
