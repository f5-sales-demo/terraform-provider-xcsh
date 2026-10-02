---
page_title: "origin_servers.origin_servers.k8s_service.site_locator"
subcategory: ""
description: "This message defines a reference to a site or virtual site object."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers origin servers k8s service site locator", "upstream servers"], "body_bytes": 2516, "body_sha256": "sha256:6b9fd15e94ecbf3e7f62c38f117fa85f3fe0d378262cf813902e29bb7d2102b4", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:site", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "site_locator", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "site_locator", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This message defines a reference to a site or virtual site object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- [origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/)
- origin_servers.origin_servers.k8s_service.site_locator

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/site/)
- [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/)
- [origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
