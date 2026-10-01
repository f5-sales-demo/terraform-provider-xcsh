---
page_title: "local_vrf.sli_config.static_v6_routes"
subcategory: ""
description: "local_vrf.sli_config.static_v6_routes for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1600, "body_sha256": "sha256:71bd66f08e003a555f3d2f55b55188db6d42762f6b674c610ebbca5a7b7467a9", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config", "path": "docs/guides/resources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf", "sli_config", "static_v6_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf.sli_config.static_v6_routes for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.sli_config.static_v6_routes

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [local_vrf](resources--securemesh_site_v2--properties--local_vrf.md)
- [local_vrf.sli_config](resources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- local_vrf.sli_config.static_v6_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [static_routes](resources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes.md): complete subsection reference.

## Next pages

- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes.md)
- [local_vrf.sli_config](resources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
