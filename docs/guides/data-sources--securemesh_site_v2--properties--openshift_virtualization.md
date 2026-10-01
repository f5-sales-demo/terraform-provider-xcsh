---
page_title: "openshift_virtualization"
subcategory: ""
description: "openshift_virtualization for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1143, "body_sha256": "sha256:6f0412834db8cbf27537944614970ae6bc215eb783a5d80b83dbb31fd152fb52", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--openshift_virtualization.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openshift_virtualization"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openshift_virtualization/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openshift_virtualization for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- openshift_virtualization

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for openshift virtualization.

Upstream description:

OpenShift Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

## Direct properties

- [not_managed](data-sources--securemesh_site_v2--properties--openshift_virtualization--not_managed.md): complete subsection reference.

## Next pages

- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--properties--openshift_virtualization--not_managed.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
