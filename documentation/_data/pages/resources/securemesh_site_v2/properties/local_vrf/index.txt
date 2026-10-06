---
page_title: "local_vrf"
subcategory: ""
description: "There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to connect LAN side"
xcsh_docs: {"aliases": ["local vrf"], "body_bytes": 2231, "body_sha256": "sha256:aa40db5a9ed28bf6ec90aa460dd0fff6cb9484e8c8327a654c1676fb11a96579", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_config", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_sli_config", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/local_vrf/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf:ConflictingObjectAttributes:default_config,slo_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf:ConflictingObjectAttributes:default_sli_config,sli_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_sli_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf:ConflictingObjectAttributes:default_sli_config,sli_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf:ConflictingObjectAttributes:default_config,slo_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf"], "schema_version": 1, "sections": [{"aliases": ["local vrf default config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "default_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf default sli config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_sli_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "default_sli_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf sli config"], "anchor": "section", "description": "Site local network configuration.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.sli_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.sli_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:no_v6_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.sli_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.sli_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes", "type": "conflicts"}], "schema_path": ["local_vrf", "sli_config"], "syntax": "block", "type": "object"}, {"aliases": ["local vrf slo config"], "anchor": "section", "description": "Site local network configuration.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:no_v6_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes", "type": "conflicts"}], "schema_path": ["local_vrf", "slo_config"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to connect LAN side", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- local_vrf

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_config",
    "slo_config"),
  validators.ConflictingObjectAttributes("default_sli_config",
    "sli_config")}
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
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

Terraform syntax:

```terraform
local_vrf {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/default_config/): complete subsection reference.

- [default_sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/default_sli_config/): complete subsection reference.

- [sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/sli_config/): complete subsection reference.

- [slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/): complete subsection reference.
