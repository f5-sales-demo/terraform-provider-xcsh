---
page_title: "code_base_integration.bitbucket_server.passwd"
subcategory: ""
description: "code_base_integration.bitbucket_server.passwd for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 2154, "body_sha256": "sha256:f1ab318cdfdd9d7c810cd8fd6f38e93a7feed706f3ec32d8e3f251b42f4ae070", "canonical_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd:blindfold_secret_info", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd:clear_secret_info"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "path": "docs/guides/resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration", "bitbucket_server", "passwd"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.bitbucket_server.passwd for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# code_base_integration.bitbucket_server.passwd

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
- [Property reference](resources--code_base_integration--reference.md)
- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- [code_base_integration.bitbucket_server](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md)
- code_base_integration.bitbucket_server.passwd

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
passwd {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--clear_secret_info.md): complete subsection reference.

## Next pages

- [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--blindfold_secret_info.md)
- [code_base_integration.bitbucket_server.passwd.clear_secret_info](resources--code_base_integration--properties--code_base_integration--bitbucket_server--passwd--clear_secret_info.md)
- [code_base_integration.bitbucket_server](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
