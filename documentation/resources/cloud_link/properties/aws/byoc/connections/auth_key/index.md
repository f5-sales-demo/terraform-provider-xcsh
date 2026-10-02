---
page_title: "aws.byoc.connections.auth_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["aws byoc connections auth key"], "body_bytes": 2497, "body_sha256": "sha256:f1f0faf360591937ee8ed3d51b947821cc215f9a6c8335a86eab4d8281c193ca", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "path": "documentation/resources/cloud_link/properties/aws/byoc/connections/auth_key/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws.byoc.connections.auth_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws.byoc.connections.auth_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc", "connections", "auth_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws--byoc--connections--auth_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "aws.byoc.connections.auth_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["aws", "byoc", "connections", "auth_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws--byoc--connections--auth_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "aws.byoc.connections.auth_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:auth_key:clear_secret_info", "type": "requires"}], "schema_path": ["aws", "byoc", "connections", "auth_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/connections/auth_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections.auth_key

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/): complete subsection reference.

## Next pages

- [aws.byoc.connections.auth_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/blindfold_secret_info/)
- [aws.byoc.connections.auth_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/auth_key/clear_secret_info/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
