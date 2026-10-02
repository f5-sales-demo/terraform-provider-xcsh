---
page_title: "local_control_plane.bgp_config"
subcategory: ""
description: "BGP configuration parameters."
xcsh_docs: {"aliases": ["local control plane bgp config"], "body_bytes": 2771, "body_sha256": "sha256:7ee62497766cf3aad7d6e0cb9e81f7dd50095bf5fee819bb45f9743a5b7e6a7b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0302110121030101-3332113333221131-1101332011331303-0023133000231111-3032313222303110-3300213220233121-0030013213102013-1133233333233201", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [{"anchor": "schema-local_control_plane--bgp_config--asn", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config:RequiredObjectAttributes:asn", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config"], "schema_version": 1, "sections": [{"aliases": ["asn"], "anchor": "schema-local_control_plane--bgp_config--asn", "description": "Autonomous System Number.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers"], "anchor": "section", "description": "BGP parameters for peer.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:bfd_disabled,bfd_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:bfd_disabled,bfd_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:bfd_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:disable_spec,routing_policies", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:ebgp_multihop_disabled,ebgp_multihop_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:ebgp_multihop_disabled,ebgp_multihop_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:ebgp_multihop_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:passive_mode_disabled,passive_mode_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:passive_mode_disabled,passive_mode_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:passive_mode_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers:ConflictingListObjectAttributes:disable_spec,routing_policies", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "type": "conflicts"}], "schema_path": ["local_control_plane", "bgp_config", "peers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BGP configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
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

- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
