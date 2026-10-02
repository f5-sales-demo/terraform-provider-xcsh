---
page_title: "local_control_plane.bgp_config.peers"
subcategory: ""
description: "BGP parameters for peer."
xcsh_docs: {"aliases": ["local control plane bgp config peers"], "body_bytes": 6719, "body_sha256": "sha256:eb1485834a53f8ce2a6a543442ede95dea171f2775accaa42f5337174132cc41", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_disabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:disable_spec", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_disabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_enabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:metadata", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_disabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_enabled", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "path": "documentation/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers"], "schema_version": 1, "sections": [{"aliases": ["bfd disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "bfd_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["bfd enabled"], "anchor": "section", "description": "BFD parameters.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "bfd_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["ebgp multihop disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "ebgp_multihop_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["ebgp multihop enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "ebgp_multihop_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["external"], "anchor": "section", "description": "External BGP Peer parameters.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external"], "syntax": "attribute", "type": "object"}, {"aliases": ["label"], "anchor": "schema-local_control_plane--bgp_config--peers--label", "description": "Specify whether this peer should be.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "label"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["passive mode disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "passive_mode_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["passive mode enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "passive_mode_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["routing policies"], "anchor": "section", "description": "List of rules which can be applied on all or particular nodes.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BGP parameters for peer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/)
- local_control_plane.bgp_config.peers

<a id="section"></a>

Type: `"list"`. Computed.

Peers. BGP parameters for peer.

Upstream description:

BGP parameters for peer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/bfd_disabled/): complete subsection reference.

- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/bfd_enabled/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/disable_spec/): complete subsection reference.

- [ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/ebgp_multihop_disabled/): complete subsection reference.

- [ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/ebgp_multihop_enabled/): complete subsection reference.

- [external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/): complete subsection reference.

<a id="schema-local_control_plane--bgp_config--peers--label"></a>

### label property

Type: `"string"`. Computed.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/metadata/): complete subsection reference.

- [passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/passive_mode_disabled/): complete subsection reference.

- [passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/passive_mode_enabled/): complete subsection reference.

- [routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/bfd_disabled/)
- [local_control_plane.bgp_config.peers.bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/bfd_enabled/)
- [local_control_plane.bgp_config.peers.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/disable_spec/)
- [local_control_plane.bgp_config.peers.ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/ebgp_multihop_disabled/)
- [local_control_plane.bgp_config.peers.ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/ebgp_multihop_enabled/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- [local_control_plane.bgp_config.peers.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/metadata/)
- [local_control_plane.bgp_config.peers.passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/passive_mode_disabled/)
- [local_control_plane.bgp_config.peers.passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/passive_mode_enabled/)
- [local_control_plane.bgp_config.peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
