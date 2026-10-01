---
page_title: "local_vrf"
subcategory: ""
description: "local_vrf for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2376, "body_sha256": "sha256:d5c9c50da2ee8ffeb8a220f02f1af0d6d25bec9eb7b1ea4e9b16262371087274", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:default_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:default_sli_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--local_vrf.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- local_vrf

<a id="section"></a>

Type: `"single"`. Computed.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF.

Upstream description:

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

- [default_config](data-sources--securemesh_site_v2--properties--local_vrf--default_config.md): complete subsection reference.

- [default_sli_config](data-sources--securemesh_site_v2--properties--local_vrf--default_sli_config.md): complete subsection reference.

- [sli_config](data-sources--securemesh_site_v2--properties--local_vrf--sli_config.md): complete subsection reference.

- [slo_config](data-sources--securemesh_site_v2--properties--local_vrf--slo_config.md): complete subsection reference.

## Next pages

- [local_vrf.default_config](data-sources--securemesh_site_v2--properties--local_vrf--default_config.md)
- [local_vrf.default_sli_config](data-sources--securemesh_site_v2--properties--local_vrf--default_sli_config.md)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--properties--local_vrf--slo_config.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
