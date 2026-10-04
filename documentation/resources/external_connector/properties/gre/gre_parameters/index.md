---
page_title: "gre.gre_parameters"
subcategory: ""
description: "GRE configuration parameters required for GRE Connection type."
xcsh_docs: {"aliases": ["gre gre parameters"], "body_bytes": 4504, "body_sha256": "sha256:cb8e40bb8316c9305289bbe9374c0ced8a5d9837b3e254699736cc2112d3cdd6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters:peer_ip_address", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "parent_id": "xcsh-docs:resources:external_connector:properties:gre", "path": "documentation/resources/external_connector/properties/gre/gre_parameters/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "type": "conflicts"}, {"anchor": "schema-gre--gre_parameters--tunnel_mtu", "enforcement": "provider-schema", "group": "gre.gre_parameters:RequiredObjectAttributes:tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:RequiredObjectAttributes:tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters"], "schema_version": 1, "sections": [{"aliases": ["gre gre parameters peer ip address"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:peer_ip_address", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre", "gre_parameters", "peer_ip_address"], "syntax": "block", "type": "object"}, {"aliases": ["gre gre parameters segment"], "anchor": "section", "description": "Reference to Segment Object.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters.segment:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment:refs", "type": "requires"}], "schema_path": ["gre", "gre_parameters", "segment"], "syntax": "block", "type": "object"}, {"aliases": ["gre gre parameters site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["gre gre parameters tunnel eps"], "anchor": "section", "description": "Configure tunnel parameters, source, destination, IP addresses.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-gre--gre_parameters--tunnel_eps--interface", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-gre--gre_parameters--tunnel_eps--node", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}], "schema_path": ["gre", "gre_parameters", "tunnel_eps"], "syntax": "block", "type": "object"}, {"aliases": ["gre gre parameters tunnel mtu"], "anchor": "schema-gre--gre_parameters--tunnel_mtu", "description": "Configure MTU for the GRE tunnel interface.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_mtu"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "GRE configuration parameters required for GRE Connection type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/)
- gre.gre_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

GRE configuration parameters required for GRE Connection type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
gre_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/peer_ip_address/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/): complete subsection reference.

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/site_local_network/): complete subsection reference.

- [tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/): complete subsection reference.

<a id="schema-gre--gre_parameters--tunnel_mtu"></a>

### tunnel_mtu property

Type: `"number"`. Optional.

Configure MTU for the GRE tunnel interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

## Next pages

- [gre.gre_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/peer_ip_address/)
- [gre.gre_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/)
- [gre.gre_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/site_local_inside_network/)
- [gre.gre_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/site_local_network/)
- [gre.gre_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
