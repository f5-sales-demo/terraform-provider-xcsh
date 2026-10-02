---
page_title: "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "List of nodes on which BGP routing policy has to be applied."
xcsh_docs: {"aliases": ["local control plane bgp config peers routing policies route policy node name"], "body_bytes": 2582, "body_sha256": "sha256:ab4453078b70ee2902cb9439dd75d65ded16cd94cd75c12e99fca31c00955626", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/node_name/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0223313210220112-3320233012323201-0221013311001211-1130301113121333-1023103312221201-0300112310101010-3301101100232331-0313113001000000", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "sections": [{"aliases": ["node"], "anchor": "schema-local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name--node", "description": "Select BGP Session on which policy will be applied.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "node_name", "node"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of nodes on which BGP routing policy has to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of nodes on which BGP routing policy has to be applied.

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
node_name {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name--node"></a>

### node property

Type: `["list", "string"]`. Optional.

Select BGP Session on which policy will be applied.

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

## Next pages

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
