---
page_title: "local_vrf"
subcategory: ""
description: "local_vrf for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2615, "body_sha256": "sha256:a2a359e632e56533c074118db0dd86d936a0407784760478be54680a1410494c", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_config", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:default_sli_config", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config", "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--local_vrf.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_vrf

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- local_vrf

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF.

Upstream description:

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Provider validators and defaults (from schema source):

```go
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

- [default_config](resources--securemesh_site_v2--properties--local_vrf--default_config.md): complete subsection reference.

- [default_sli_config](resources--securemesh_site_v2--properties--local_vrf--default_sli_config.md): complete subsection reference.

- [sli_config](resources--securemesh_site_v2--properties--local_vrf--sli_config.md): complete subsection reference.

- [slo_config](resources--securemesh_site_v2--properties--local_vrf--slo_config.md): complete subsection reference.

## Next pages

- [local_vrf.default_config](resources--securemesh_site_v2--properties--local_vrf--default_config.md)
- [local_vrf.default_sli_config](resources--securemesh_site_v2--properties--local_vrf--default_sli_config.md)
- [local_vrf.sli_config](resources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- [local_vrf.slo_config](resources--securemesh_site_v2--properties--local_vrf--slo_config.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
