---
page_title: "where.site.disable_internet_vip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["where site disable internet vip"], "body_bytes": 1282, "body_sha256": "sha256:aa0140bb939437d5728f7d7ea66e0f86d103eb545ff7a93449b5c8ab0cccdef7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where:site:disable_internet_vip", "parent_id": "xcsh-docs:resources:bgp:properties:where:site", "path": "documentation/resources/bgp/properties/where/site/disable_internet_vip/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2010120022210101-0231321322203131-0323101320233313-3031000003210202-2032021131102102-3220021302000223-3020301013303332-3300300032103102", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "site", "disable_internet_vip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/site/disable_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.site.disable_internet_vip

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/)
- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/site/)
- where.site.disable_internet_vip

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
disable_internet_vip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/site/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
