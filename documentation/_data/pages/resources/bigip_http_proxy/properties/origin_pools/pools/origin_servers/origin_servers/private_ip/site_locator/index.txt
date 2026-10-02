---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator"
subcategory: ""
description: "This message defines a reference to a site or virtual site object."
xcsh_docs: {"aliases": ["backend servers", "origin pools pools origin servers origin servers private ip site locator", "origin servers", "upstream servers"], "body_bytes": 3417, "body_sha256": "sha256:cea994e450a3cc67f6c45b7447f013c1082d544c11196eeb8e07ef6f86c5b28e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:site", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1230023122312032-1103223311003112-2213211131230113-1201222200312112-0222310213120110-1103130022310122-2131113221210111-2032330032013313", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "site_locator"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--name", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:site", "type": "requires"}], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "site_locator", "site"], "syntax": "block", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--name", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:virtual_site", "type": "requires"}], "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "site_locator", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This message defines a reference to a site or virtual site object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
```

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

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/virtual_site/): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/site/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/virtual_site/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
