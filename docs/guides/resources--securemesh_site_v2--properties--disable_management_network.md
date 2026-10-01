---
page_title: "disable_management_network"
subcategory: ""
description: "disable_management_network for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1424, "body_sha256": "sha256:ecd1e12ea7f80ce57cb48d7cc76bac09ca6bc2a36fc8371aa79e62043e948fa2", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_management_network", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_management_network", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--disable_management_network.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_management_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/disable_management_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_management_network for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_management_network

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- disable_management_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_management\_network, enable\_management\_network; Default:
disable\_management\_network\] Configuration parameter for disable management network.

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

- [disable_management_network](resources--securemesh_site_v2--properties--disable_management_network.md#section)
- [enable_management_network](resources--securemesh_site_v2--properties--enable_management_network.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_management_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
