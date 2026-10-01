---
page_title: "aws.byoc.connections.auth_key"
subcategory: ""
description: "aws.byoc.connections.auth_key for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1955, "body_sha256": "sha256:3d299610f935fff701abbe84a86e3b12abce807c4abe49cc8e682287d5269485", "canonical_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "path": "docs/guides/resources--cloud_link--properties--aws--byoc--connections--auth_key.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws", "byoc", "connections", "auth_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/connections/auth_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws.byoc.connections.auth_key for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections.auth_key

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md)
- [Property reference](resources--cloud_link--reference.md)
- [aws](resources--cloud_link--properties--aws.md)
- [aws.byoc](resources--cloud_link--properties--aws--byoc.md)
- [aws.byoc.connections](resources--cloud_link--properties--aws--byoc--connections.md)
- aws.byoc.connections.auth_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
auth_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--cloud_link--properties--aws--byoc--connections--auth_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--cloud_link--properties--aws--byoc--connections--auth_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [aws.byoc.connections.auth_key.blindfold_secret_info](resources--cloud_link--properties--aws--byoc--connections--auth_key--blindfold_secret_info.md)
- [aws.byoc.connections.auth_key.clear_secret_info](resources--cloud_link--properties--aws--byoc--connections--auth_key--clear_secret_info.md)
- [aws.byoc.connections](resources--cloud_link--properties--aws--byoc--connections.md)
- [xcsh_cloud_link](../resources/cloud_link.md)
