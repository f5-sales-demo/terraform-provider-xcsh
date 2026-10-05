---
page_title: "local_control_plane.bgp_config.peers.bfd_enabled"
subcategory: ""
description: "BFD parameters."
xcsh_docs: {"aliases": ["local control plane bgp config peers bfd enabled"], "body_bytes": 5428, "body_sha256": "sha256:81524cdc2c18d3c824a4b6302a0a9ce545330b3864dfe09a3fcd7faba6261c13", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/bfd_enabled/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1020332033000232-0000011311000002-2023210202110001-3301131011123332-1010132232200221-0023222102323121-2103300213013222-3212323112323321", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [{"anchor": "schema-local_control_plane--bgp_config--peers--bfd_enabled--multiplier", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers.bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "type": "requires"}, {"anchor": "schema-local_control_plane--bgp_config--peers--bfd_enabled--receive_interval_milliseconds", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers.bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "type": "requires"}, {"anchor": "schema-local_control_plane--bgp_config--peers--bfd_enabled--transmit_interval_milliseconds", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers.bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "bfd_enabled"], "schema_version": 1, "sections": [{"aliases": ["local control plane bgp config peers bfd enabled multiplier"], "anchor": "schema-local_control_plane--bgp_config--peers--bfd_enabled--multiplier", "description": "Specify Number of missed packets to bring session down\"", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "bfd_enabled", "multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["local control plane bgp config peers bfd enabled receive interval milliseconds"], "anchor": "schema-local_control_plane--bgp_config--peers--bfd_enabled--receive_interval_milliseconds", "description": "BFD receive interval timer, in milliseconds.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "bfd_enabled", "receive_interval_milliseconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["local control plane bgp config peers bfd enabled transmit interval milliseconds"], "anchor": "schema-local_control_plane--bgp_config--peers--bfd_enabled--transmit_interval_milliseconds", "description": "BFD transmit interval timer, in milliseconds.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "bfd_enabled", "transmit_interval_milliseconds"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/bfd_enabled/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "BFD parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.bfd_enabled

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- local_control_plane.bgp_config.peers.bfd_enabled

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Upstream description:

BFD parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
```

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
bfd_enabled {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_control_plane--bgp_config--peers--bfd_enabled--multiplier"></a>

### multiplier property

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="schema-local_control_plane--bgp_config--peers--bfd_enabled--receive_interval_milliseconds"></a>

### receive_interval_milliseconds property

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="schema-local_control_plane--bgp_config--peers--bfd_enabled--transmit_interval_milliseconds"></a>

### transmit_interval_milliseconds property

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

## Next pages

- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
