---
page_title: "peers"
subcategory: ""
description: "peers for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 5143, "body_sha256": "sha256:897b5d174d87631acbf253101468c998e408f1ef77d26139a65a7408facbbf7f", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:bfd_disabled", "xcsh-docs:data-sources:bgp:properties:peers:bfd_enabled", "xcsh-docs:data-sources:bgp:properties:peers:disable_spec", "xcsh-docs:data-sources:bgp:properties:peers:ebgp_multihop_disabled", "xcsh-docs:data-sources:bgp:properties:peers:ebgp_multihop_enabled", "xcsh-docs:data-sources:bgp:properties:peers:external", "xcsh-docs:data-sources:bgp:properties:peers:metadata", "xcsh-docs:data-sources:bgp:properties:peers:passive_mode_disabled", "xcsh-docs:data-sources:bgp:properties:peers:passive_mode_enabled", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers", "parent_id": "xcsh-docs:data-sources:bgp:reference", "path": "documentation/data-sources/bgp/properties/peers/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["peers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- peers

<a id="section"></a>

Type: `"list"`. Computed.

Peers. List of peers.

Upstream description:

List of peers.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_disabled/): complete subsection reference.

- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_enabled/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/disable_spec/): complete subsection reference.

- [ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/ebgp_multihop_disabled/): complete subsection reference.

- [ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/ebgp_multihop_enabled/): complete subsection reference.

- [external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/): complete subsection reference.

<a id="schema-peers--label"></a>

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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/metadata/): complete subsection reference.

- [passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/passive_mode_disabled/): complete subsection reference.

- [passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/passive_mode_enabled/): complete subsection reference.

- [routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/): complete subsection reference.

## Next pages

- [peers.bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_disabled/)
- [peers.bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_enabled/)
- [peers.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/disable_spec/)
- [peers.ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/ebgp_multihop_disabled/)
- [peers.ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/ebgp_multihop_enabled/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- [peers.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/metadata/)
- [peers.passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/passive_mode_disabled/)
- [peers.passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/passive_mode_enabled/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
