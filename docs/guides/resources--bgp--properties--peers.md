---
page_title: "peers"
subcategory: ""
description: "peers for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 4438, "body_sha256": "sha256:23f271d449d2dfaa13a8d26c1bea03a379104d4582287e1760382660f7b26ac5", "canonical_id": "xcsh-docs:resources:bgp:properties:peers", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:bfd_disabled", "xcsh-docs:resources:bgp:properties:peers:bfd_enabled", "xcsh-docs:resources:bgp:properties:peers:disable_spec", "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_disabled", "xcsh-docs:resources:bgp:properties:peers:ebgp_multihop_enabled", "xcsh-docs:resources:bgp:properties:peers:external", "xcsh-docs:resources:bgp:properties:peers:metadata", "xcsh-docs:resources:bgp:properties:peers:passive_mode_disabled", "xcsh-docs:resources:bgp:properties:peers:passive_mode_enabled", "xcsh-docs:resources:bgp:properties:peers:routing_policies"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "docs/guides/resources--bgp--properties--peers.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
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

- [bfd_disabled](resources--bgp--properties--peers--bfd_disabled.md): complete subsection reference.

- [bfd_enabled](resources--bgp--properties--peers--bfd_enabled.md): complete subsection reference.

- [disable_spec](resources--bgp--properties--peers--disable_spec.md): complete subsection reference.

- [ebgp_multihop_disabled](resources--bgp--properties--peers--ebgp_multihop_disabled.md): complete subsection reference.

- [ebgp_multihop_enabled](resources--bgp--properties--peers--ebgp_multihop_enabled.md): complete subsection reference.

- [external](resources--bgp--properties--peers--external.md): complete subsection reference.

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

- [metadata](resources--bgp--properties--peers--metadata.md): complete subsection reference.

- [passive_mode_disabled](resources--bgp--properties--peers--passive_mode_disabled.md): complete subsection reference.

- [passive_mode_enabled](resources--bgp--properties--peers--passive_mode_enabled.md): complete subsection reference.

- [routing_policies](resources--bgp--properties--peers--routing_policies.md): complete subsection reference.

## Next pages

- [peers.bfd_disabled](resources--bgp--properties--peers--bfd_disabled.md)
- [peers.bfd_enabled](resources--bgp--properties--peers--bfd_enabled.md)
- [peers.disable_spec](resources--bgp--properties--peers--disable_spec.md)
- [peers.ebgp_multihop_disabled](resources--bgp--properties--peers--ebgp_multihop_disabled.md)
- [peers.ebgp_multihop_enabled](resources--bgp--properties--peers--ebgp_multihop_enabled.md)
- [peers.external](resources--bgp--properties--peers--external.md)
- [peers.metadata](resources--bgp--properties--peers--metadata.md)
- [peers.passive_mode_disabled](resources--bgp--properties--peers--passive_mode_disabled.md)
- [peers.passive_mode_enabled](resources--bgp--properties--peers--passive_mode_enabled.md)
- [peers.routing_policies](resources--bgp--properties--peers--routing_policies.md)
- [Property reference](resources--bgp--reference.md)
- [xcsh_bgp](../resources/bgp.md)
