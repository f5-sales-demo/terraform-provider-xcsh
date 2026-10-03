---
page_title: "segment_vrf"
subcategory: ""
description: "The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name. Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per Site that can be configured here."
xcsh_docs: {"aliases": ["segment vrf"], "body_bytes": 2592, "body_sha256": "sha256:fc9813589ab9bdbc6f07aa2848d216fd5ec36d0619e10b0fbe0197f559f55dd3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config", "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_network"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/segment_vrf/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf"], "schema_version": 1, "sections": [{"aliases": ["segment vrf segment config"], "anchor": "section", "description": "Segment Network Configuration.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:no_v6_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes", "type": "conflicts"}], "schema_path": ["segment_vrf", "segment_config"], "syntax": "block", "type": "object"}, {"aliases": ["segment vrf segment network"], "anchor": "section", "description": "This type establishes a 'direct reference' from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name for public API and Uid for private API This type of reference is called direct because the relation is explicit and concrete (as opposed to selector reference which", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["segment_vrf", "segment_network"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/segment_vrf/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name. Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per Site that can be configured here.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- segment_vrf

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

Upstream description:

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
segment_vrf {
  # Configure direct properties listed below.
}
```

## Direct properties

- [segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/): complete subsection reference.

- [segment_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_network/): complete subsection reference.

## Next pages

- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- [segment_vrf.segment_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
