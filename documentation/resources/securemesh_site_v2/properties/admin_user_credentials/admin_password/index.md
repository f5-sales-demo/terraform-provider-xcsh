---
page_title: "admin_user_credentials.admin_password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["admin user credentials admin password", "authentication", "credential setup", "credentials"], "body_bytes": 2449, "body_sha256": "sha256:8e739354371e7c03955f0286b6f901e3f8af5ecdb49bb96147108e8aef8d73d8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:blindfold_secret_info", "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:clear_secret_info"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials", "path": "documentation/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1221322130013030-0212023021201321-0232021201030201-2301320123003221-1210230200331322-0003332110201133-3203212221321213-1320113022233000", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "admin_user_credentials.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "admin_user_credentials.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["admin_user_credentials", "admin_password"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-admin_user_credentials--admin_password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "admin_user_credentials.admin_password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:blindfold_secret_info", "type": "requires"}], "schema_path": ["admin_user_credentials", "admin_password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-admin_user_credentials--admin_password--clear_secret_info--url", "enforcement": "provider-schema", "group": "admin_user_credentials.admin_password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:clear_secret_info", "type": "requires"}], "schema_path": ["admin_user_credentials", "admin_password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_user_credentials.admin_password

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [admin_user_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/)
- admin_user_credentials.admin_password

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
admin_password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/clear_secret_info/): complete subsection reference.

## Next pages

- [admin_user_credentials.admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/blindfold_secret_info/)
- [admin_user_credentials.admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/clear_secret_info/)
- [admin_user_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
