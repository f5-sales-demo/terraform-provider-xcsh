---
page_title: "site_mesh_group_on_slo"
subcategory: ""
description: "site_mesh_group_on_slo for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2569, "body_sha256": "sha256:7ae776a43c0e150d5d23f3c65327cb1d805614e2ab1fc584c4523f0302e94070", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--site_mesh_group_on_slo.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_mesh_group_on_slo"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_mesh_group_on_slo for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_mesh_group_on_slo

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- site_mesh_group_on_slo

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_site_mesh_group",
    "site_mesh_group"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
site_mesh_group_on_slo {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_site_mesh_group](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--no_site_mesh_group.md): complete subsection reference.

- [site_mesh_group](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--site_mesh_group.md): complete subsection reference.

- [sm_connection_public_ip](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--sm_connection_pvt_ip.md): complete subsection reference.

## Next pages

- [site_mesh_group_on_slo.no_site_mesh_group](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--no_site_mesh_group.md)
- [site_mesh_group_on_slo.site_mesh_group](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--site_mesh_group.md)
- [site_mesh_group_on_slo.sm_connection_public_ip](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--sm_connection_public_ip.md)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](resources--securemesh_site_v2--properties--site_mesh_group_on_slo--sm_connection_pvt_ip.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
