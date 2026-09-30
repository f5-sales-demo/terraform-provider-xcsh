---
page_title: "aws.byoc"
subcategory: ""
description: "aws.byoc for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1496, "body_sha256": "sha256:29e63dcd29a3ee3465d0c886e99b50bd7c9ef4ede88f82f2e182163dd5994bd9", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws:byoc:connections"], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws", "path": "documentation/resources/cloud_link/properties/aws/byoc/index.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["aws", "byoc"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws.byoc for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws.byoc

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
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

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/): complete subsection reference.

## Next pages

- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
