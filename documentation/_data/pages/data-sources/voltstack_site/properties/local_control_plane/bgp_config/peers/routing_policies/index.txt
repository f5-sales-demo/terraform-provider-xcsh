---
page_title: "local_control_plane.bgp_config.peers.routing_policies"
subcategory: ""
description: "List of rules which can be applied on all or particular nodes."
xcsh_docs: {"aliases": ["local control plane bgp config peers routing policies"], "body_bytes": 1997, "body_sha256": "sha256:0d1a58a5f15cf23997d9448a9a549f8f6397059f020c08e06bf4a4140e2293b0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "path": "documentation/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies"], "schema_version": 1, "sections": [{"aliases": ["route policy"], "anchor": "section", "description": "Route policy to be applied.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of rules which can be applied on all or particular nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.routing_policies

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- local_control_plane.bgp_config.peers.routing_policies

<a id="section"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

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

- [route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
