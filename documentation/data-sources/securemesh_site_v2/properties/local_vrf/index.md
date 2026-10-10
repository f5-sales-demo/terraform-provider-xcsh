---
page_title: "local_vrf"
subcategory: ""
description: "There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to connect LAN side"
xcsh_docs: {"aliases": ["local vrf"], "body_bytes": 1845, "body_sha256": "sha256:619fe0ed46acdbab51795c23c67cb1c45dbc477ecbf7d822964e838d9cbed62e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:default_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:default_sli_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/local_vrf/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf"], "schema_version": 1, "sections": [{"aliases": ["local vrf default config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:default_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "default_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf default sli config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:default_sli_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "default_sli_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf sli config"], "anchor": "section", "description": "Site local network configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_vrf", "sli_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf slo config"], "anchor": "section", "description": "Site local network configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_vrf", "slo_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to connect LAN side", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- local_vrf

<a id="section"></a>

Type: `"single"`. Computed.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

## Direct properties

- [default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/default_config/): complete subsection reference.

- [default_sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/default_sli_config/): complete subsection reference.

- [sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/): complete subsection reference.

- [slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/): complete subsection reference.
