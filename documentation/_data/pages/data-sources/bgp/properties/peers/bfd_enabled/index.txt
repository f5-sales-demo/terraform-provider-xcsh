---
page_title: "peers.bfd_enabled"
subcategory: ""
description: "peers.bfd_enabled for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 3930, "body_sha256": "sha256:66ee16b67153c251a539cafc3d004014f900cf097afc0af9965c2658fb26f3fd", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:bfd_enabled", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers", "path": "documentation/data-sources/bgp/properties/peers/bfd_enabled/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["peers", "bfd_enabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/bfd_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.bfd_enabled for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# peers.bfd_enabled

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- peers.bfd_enabled

<a id="section"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

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

## Direct properties

<a id="schema-peers--bfd_enabled--multiplier"></a>

### multiplier property

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-peers--bfd_enabled--receive_interval_milliseconds"></a>

### receive_interval_milliseconds property

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-peers--bfd_enabled--transmit_interval_milliseconds"></a>

### transmit_interval_milliseconds property

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
