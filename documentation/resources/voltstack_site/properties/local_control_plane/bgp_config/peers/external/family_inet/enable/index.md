---
page_title: "local_control_plane.bgp_config.peers.external.family_inet.enable"
subcategory: ""
description: "IPv4 Unicast."
xcsh_docs: {"aliases": ["local control plane bgp config peers external family inet enable"], "body_bytes": 2585, "body_sha256": "sha256:498aa3674efb0a8adaddb42773f7af3a224d4e12a72112f52ebbc6e65068f238", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0200101313301210-3003002131302001-3101012211103232-3322223332121013-1002322020002303-2113002323311020-2133233032123302-0210001323202302", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable"], "schema_version": 1, "sections": [{"aliases": ["aggregation"], "anchor": "section", "description": "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable:aggregation", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable", "aggregation"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPv4 Unicast.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet.enable

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- [local_control_plane.bgp_config.peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/)
- local_control_plane.bgp_config.peers.external.family_inet.enable

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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
enable {
  # Configure direct properties listed below.
}
```

## Direct properties

- [aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/aggregation/)
- [local_control_plane.bgp_config.peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
