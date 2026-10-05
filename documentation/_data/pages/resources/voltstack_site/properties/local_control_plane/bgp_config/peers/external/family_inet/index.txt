---
page_title: "local_control_plane.bgp_config.peers.external.family_inet"
subcategory: ""
description: "Parameters for inet family."
xcsh_docs: {"aliases": ["local control plane bgp config peers external family inet"], "body_bytes": 3012, "body_sha256": "sha256:6836ba72676bfee6c7e0d1c01492eb1a7dc75f67ec86991ad13d6eecc9989757", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:disable_spec", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0131120213031302-3100221121201101-2122211011020203-1000312230103002-2301302220003200-1023222122031202-0010000113121310-0100321223013011", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers.external.family_inet:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers.external.family_inet:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet"], "schema_version": 1, "sections": [{"aliases": ["local control plane bgp config peers external family inet disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["local control plane bgp config peers external family inet enable"], "anchor": "section", "description": "IPv4 Unicast.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet:enable", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "family_inet", "enable"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Parameters for inet family.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.family_inet

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- local_control_plane.bgp_config.peers.external.family_inet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
family_inet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.external.family_inet.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/disable_spec/)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/enable/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
