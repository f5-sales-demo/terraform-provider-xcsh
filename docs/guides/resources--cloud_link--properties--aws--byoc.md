---
page_title: "aws.byoc"
subcategory: ""
description: "aws.byoc for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1141, "body_sha256": "sha256:9632ff5f88c0bada771c97bfb174524c54d4806e53da95e5fce04acbdd38c8c6", "canonical_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws:byoc:connections"], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws", "path": "docs/guides/resources--cloud_link--properties--aws--byoc.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws", "byoc"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws.byoc for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws.byoc

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md)
- [Property reference](resources--cloud_link--reference.md)
- [aws](resources--cloud_link--properties--aws.md)
- aws.byoc

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bring Your Own Connections. List of Bring You Own Connection.

Upstream description:

List of Bring You Own Connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
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
byoc {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connections](resources--cloud_link--properties--aws--byoc--connections.md): complete subsection reference.

## Next pages

- [aws.byoc.connections](resources--cloud_link--properties--aws--byoc--connections.md)
- [aws](resources--cloud_link--properties--aws.md)
- [xcsh_cloud_link](../resources/cloud_link.md)
