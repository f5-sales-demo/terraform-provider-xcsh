---
page_title: "gcp.byoc.connections.same_as_credential"
subcategory: ""
description: "gcp.byoc.connections.same_as_credential for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1175, "body_sha256": "sha256:bff132a19458a6610392dd4a28afb4d107e57338bde86a7c3c14269defba4fb2", "canonical_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential", "parent_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "path": "docs/guides/resources--cloud_link--properties--gcp--byoc--connections--same_as_credential.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp", "byoc", "connections", "same_as_credential"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/gcp/byoc/connections/same_as_credential/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp.byoc.connections.same_as_credential for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.byoc.connections.same_as_credential

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md)
- [Property reference](resources--cloud_link--reference.md)
- [gcp](resources--cloud_link--properties--gcp.md)
- [gcp.byoc](resources--cloud_link--properties--gcp--byoc.md)
- [gcp.byoc.connections](resources--cloud_link--properties--gcp--byoc--connections.md)
- gcp.byoc.connections.same_as_credential

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as credential.

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

Terraform syntax:

```terraform
same_as_credential = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [gcp.byoc.connections](resources--cloud_link--properties--gcp--byoc--connections.md)
- [xcsh_cloud_link](../resources/cloud_link.md)
