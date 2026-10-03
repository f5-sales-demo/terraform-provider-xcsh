---
page_title: "local_control_plane"
subcategory: ""
description: "Enable local control plane for L3VPN, SRV6, EVPN etc."
xcsh_docs: {"aliases": ["local control plane"], "body_bytes": 2867, "body_sha256": "sha256:ef3717af50a8dd48971d05d24a6ffc6237ee8dd44e1073daf9e55745373187d8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:inside_vn", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:outside_vn"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/local_control_plane/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane:ConflictingObjectAttributes:inside_vn,outside_vn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:inside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_control_plane:ConflictingObjectAttributes:inside_vn,outside_vn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:outside_vn", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane"], "schema_version": 1, "sections": [{"aliases": ["local control plane bgp config"], "anchor": "section", "description": "BGP configuration parameters.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-local_control_plane--bgp_config--asn", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config:RequiredObjectAttributes:asn", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config", "type": "requires"}], "schema_path": ["local_control_plane", "bgp_config"], "syntax": "block", "type": "object"}, {"aliases": ["local control plane inside vn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:inside_vn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "inside_vn"], "syntax": "attribute", "type": "object"}, {"aliases": ["local control plane outside vn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:outside_vn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "outside_vn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Enable local control plane for L3VPN, SRV6, EVPN etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- local_control_plane

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: local\_control\_plane, no\_local\_control\_plane; Default: no\_local\_control\_plane\]
Enable local control plane for L3VPN, SRV6, EVPN etc.

Upstream description:

Enable local control plane for L3VPN, SRV6, EVPN etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_vn",
    "outside_vn")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_vn\",\"outside_vn\"]"
}
```

OneOf alternatives in this subsection:

- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/#section)
- [no_local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/no_local_control_plane/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
local_control_plane {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/): complete subsection reference.

- [inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/inside_vn/): complete subsection reference.

- [outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/outside_vn/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/inside_vn/)
- [local_control_plane.outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/outside_vn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
