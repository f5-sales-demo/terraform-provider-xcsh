---
page_title: "disable_ssh_access"
subcategory: ""
description: "disable_ssh_access for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1240, "body_sha256": "sha256:eba15a0866ee64ed7c2b7d92b95ff6f869be28bc7fa7fc62f4f7fd3a581aec41", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:disable_ssh_access", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:disable_ssh_access", "parent_id": "xcsh-docs:data-sources:nfv_service:reference", "path": "docs/guides/data-sources--nfv_service--properties--disable_ssh_access.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_ssh_access"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/disable_ssh_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_ssh_access for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ssh_access

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- disable_ssh_access

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ssh\_access, enabled\_ssh\_access; Default: disable\_ssh\_access\] Configuration
parameter for disable ssh access.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [disable_ssh_access](data-sources--nfv_service--properties--disable_ssh_access.md#section)
- [enabled_ssh_access](data-sources--nfv_service--properties--enabled_ssh_access.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--nfv_service--reference.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
