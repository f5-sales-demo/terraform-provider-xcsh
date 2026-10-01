---
page_title: "peers"
subcategory: ""
description: "peers for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 5646, "body_sha256": "sha256:3e37de02264f487f47b6ca37947ec9366dcfe7aa8aca6295bf234fbefb5cd6a5", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:bfd_disabled", "xcsh-docs:resources:bgp:properties:peers:bfd_enabled", "xcsh-docs:resources:bgp:properties:peers:disable_spec", "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_disabled", "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_enabled", "xcsh-docs:resources:bgp:properties:peers:external", "xcsh-docs:resources:bgp:properties:peers:metadata", "xcsh-docs:resources:bgp:properties:peers:passive_mode_disabled", "xcsh-docs:resources:bgp:properties:peers:passive_mode_enabled", "xcsh-docs:resources:bgp:properties:peers:routing_policies"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "documentation/resources/bgp/properties/peers/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["peers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- peers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Peers. List of peers.

Upstream description:

List of peers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bfd_disabled",
    "bfd_enabled"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "routing_policies"),
  validators.ConflictingListObjectAttributes("ebgp_multihop_disabled",
    "ebgp_multihop_enabled"),
  validators.ConflictingListObjectAttributes("passive_mode_disabled",
    "passive_mode_enabled")}
```

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

Terraform syntax:

```terraform
peers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/bfd_disabled/): complete subsection reference.

- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/bfd_enabled/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/disable_spec/): complete subsection reference.

- [ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/ebgp_multihop_disabled/): complete subsection reference.

- [ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/ebgp_multihop_enabled/): complete subsection reference.

- [external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/): complete subsection reference.

<a id="schema-peers--label"></a>

### label property

Type: `"string"`. Optional.

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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/metadata/): complete subsection reference.

- [passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/passive_mode_disabled/): complete subsection reference.

- [passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/passive_mode_enabled/): complete subsection reference.

- [routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/): complete subsection reference.

## Next pages

- [peers.bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/bfd_disabled/)
- [peers.bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/bfd_enabled/)
- [peers.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/disable_spec/)
- [peers.ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/ebgp_multihop_disabled/)
- [peers.ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/ebgp_multihop_enabled/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- [peers.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/metadata/)
- [peers.passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/passive_mode_disabled/)
- [peers.passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/passive_mode_enabled/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
