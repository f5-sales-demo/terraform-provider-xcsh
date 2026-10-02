---
page_title: "peers.routing_policies.route_policy"
subcategory: ""
description: "Route policy to be applied."
xcsh_docs: {"aliases": ["peers routing policies route policy"], "body_bytes": 3826, "body_sha256": "sha256:b62578dbfc39c179c89b50a65aa85b7ed722803b319f0ea8dd6a79046a33ac6c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:object_refs", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies", "path": "documentation/resources/bgp/properties/peers/routing_policies/route_policy/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:all_nodes,node_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:inbound,outbound", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:all_nodes,node_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:inbound,outbound", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:RequiredListObjectAttributes:object_refs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:object_refs", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy"], "schema_version": 1, "sections": [{"aliases": ["all nodes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "all_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["inbound"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "inbound"], "syntax": "attribute", "type": "object"}, {"aliases": ["node name"], "anchor": "section", "description": "List of nodes on which BGP routing policy has to be applied.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "node_name"], "syntax": "block", "type": "object"}, {"aliases": ["object refs"], "anchor": "section", "description": "Select route policy to apply.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:object_refs", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "object_refs"], "syntax": "block", "type": "object"}, {"aliases": ["outbound"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "outbound"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Route policy to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgpCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/)
- peers.routing_policies.route_policy

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

- [all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/all_nodes/): complete subsection reference.

- [inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/inbound/): complete subsection reference.

- [node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/node_name/): complete subsection reference.

- [object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/object_refs/): complete subsection reference.

- [outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/outbound/): complete subsection reference.

## Next pages

- [peers.routing_policies.route_policy.all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/all_nodes/)
- [peers.routing_policies.route_policy.inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/inbound/)
- [peers.routing_policies.route_policy.node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/node_name/)
- [peers.routing_policies.route_policy.object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/object_refs/)
- [peers.routing_policies.route_policy.outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/outbound/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
