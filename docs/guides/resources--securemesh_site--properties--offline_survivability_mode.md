---
page_title: "offline_survivability_mode"
subcategory: ""
description: "offline_survivability_mode for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 2741, "body_sha256": "sha256:baa23152f8987ce9052f9c438e3641f31ca5f24a2a466c7a97684f94d5c5b638", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode:no_offline_survivability_mode"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode", "parent_id": "xcsh-docs:resources:securemesh_site:reference", "path": "docs/guides/resources--securemesh_site--properties--offline_survivability_mode.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "offline_survivability_mode for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- offline_survivability_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enable_offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md): complete subsection reference.

- [no_offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode--no_offline_survivability_mode.md): complete subsection reference.

## Next pages

- [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md)
- [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode--no_offline_survivability_mode.md)
- [Property reference](resources--securemesh_site--reference.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
