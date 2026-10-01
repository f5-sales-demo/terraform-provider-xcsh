---
page_title: "local_control_plane.bgp_config"
subcategory: ""
description: "local_control_plane.bgp_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2416, "body_sha256": "sha256:c64505d5a42de9f709c07a603f0cc2e612914dec9c8e83f672b4b31f732c486f", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- local_control_plane.bgp_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BGP Configuration. BGP configuration parameters.

Upstream description:

BGP configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn")}
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
bgp_config {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_control_plane--bgp_config--asn"></a>

### asn property

Type: `"number"`. Optional.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
