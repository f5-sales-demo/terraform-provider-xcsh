---
page_title: "disable_ssh_access"
subcategory: ""
description: "disable_ssh_access for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1283, "body_sha256": "sha256:2f2adfd594c512feb2a604e639f75491848119ae1b357efd7a0437eb3b8bb8f9", "canonical_id": "xcsh-docs:resources:nfv_service:properties:disable_ssh_access", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:disable_ssh_access", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "docs/guides/resources--nfv_service--properties--disable_ssh_access.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_ssh_access"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/disable_ssh_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_ssh_access for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ssh_access

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- disable_ssh_access

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

- [disable_ssh_access](resources--nfv_service--properties--disable_ssh_access.md#section)
- [enabled_ssh_access](resources--nfv_service--properties--enabled_ssh_access.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ssh_access = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--nfv_service--reference.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
