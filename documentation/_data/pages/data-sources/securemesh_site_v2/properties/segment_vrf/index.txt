---
page_title: "segment_vrf"
subcategory: ""
description: "The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name. Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per Site that can be configured here."
xcsh_docs: {"aliases": ["segment vrf"], "body_bytes": 1662, "body_sha256": "sha256:2c5b7e81da54e16f07499fb54f1dc36f69e9a5bc15b95f74901f683ebfa57985", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_network"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/segment_vrf/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf"], "schema_version": 1, "sections": [{"aliases": ["segment vrf segment config"], "anchor": "section", "description": "Segment Network Configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["segment_vrf", "segment_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["segment vrf segment network"], "anchor": "section", "description": "This type establishes a 'direct reference' from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name for public API and Uid for private API This type of reference is called direct because the relation is explicit and concrete (as opposed to selector reference which", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["segment_vrf", "segment_network"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/segment_vrf/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name. Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per Site that can be configured here.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- segment_vrf

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

## Direct properties

- [segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/): complete subsection reference.

- [segment_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_network/): complete subsection reference.
