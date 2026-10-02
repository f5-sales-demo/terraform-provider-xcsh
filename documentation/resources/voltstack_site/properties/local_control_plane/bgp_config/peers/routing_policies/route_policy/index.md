---
page_title: "local_control_plane.bgp_config.peers.routing_policies.route_policy"
subcategory: ""
description: "Route policy to be applied."
xcsh_docs: {"aliases": ["local control plane bgp config peers routing policies route policy"], "body_bytes": 5035, "body_sha256": "sha256:449009395d825d3cab42aacb3f1ae50a24bba7cb84312eadb955d39dbe13e2c5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:inbound", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:object_refs", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:outbound"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy"], "schema_version": 1, "sections": [{"aliases": ["all nodes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:all_nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "all_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["inbound"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:inbound", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "inbound"], "syntax": "attribute", "type": "object"}, {"aliases": ["node name"], "anchor": "section", "description": "List of nodes on which BGP routing policy has to be applied.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "node_name"], "syntax": "block", "type": "object"}, {"aliases": ["object refs"], "anchor": "section", "description": "Select route policy to apply.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:object_refs", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "object_refs"], "syntax": "block", "type": "object"}, {"aliases": ["outbound"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:outbound", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "outbound"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Route policy to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.routing_policies.route_policy

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/)
- local_control_plane.bgp_config.peers.routing_policies.route_policy

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("object_refs"),
  validators.ConflictingListObjectAttributes("all_nodes",
    "node_name"),
  validators.ConflictingListObjectAttributes("inbound",
    "outbound")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
route_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/all_nodes/): complete subsection reference.

- [inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/inbound/): complete subsection reference.

- [node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/node_name/): complete subsection reference.

- [object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/object_refs/): complete subsection reference.

- [outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/outbound/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/all_nodes/)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/inbound/)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/node_name/)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/object_refs/)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/outbound/)
- [local_control_plane.bgp_config.peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
