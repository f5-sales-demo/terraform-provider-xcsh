---
page_title: "domains.simple_login.password"
subcategory: ""
description: "domains.simple_login.password for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 1907, "body_sha256": "sha256:5f6460fc0145ca7c2a2fd84db1181b62c6f46c078f049256c65510566b16aa3e", "canonical_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password", "child_ids": ["xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password", "parent_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login", "path": "docs/guides/resources--api_crawler--properties--domains--simple_login--password.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "simple_login", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/properties/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.simple_login.password for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.simple_login.password

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md)
- [Property reference](resources--api_crawler--reference.md)
- [domains](resources--api_crawler--properties--domains.md)
- [domains.simple_login](resources--api_crawler--properties--domains--simple_login.md)
- domains.simple_login.password

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
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--api_crawler--properties--domains--simple_login--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--api_crawler--properties--domains--simple_login--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [domains.simple_login.password.blindfold_secret_info](resources--api_crawler--properties--domains--simple_login--password--blindfold_secret_info.md)
- [domains.simple_login.password.clear_secret_info](resources--api_crawler--properties--domains--simple_login--password--clear_secret_info.md)
- [domains.simple_login](resources--api_crawler--properties--domains--simple_login.md)
- [xcsh_api_crawler](../resources/api_crawler.md)
