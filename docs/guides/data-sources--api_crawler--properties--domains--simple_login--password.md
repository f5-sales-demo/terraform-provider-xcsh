---
page_title: "domains.simple_login.password"
subcategory: ""
description: "domains.simple_login.password for xcsh_api_crawler."
xcsh_docs: {"aliases": [], "body_bytes": 1533, "body_sha256": "sha256:280f1be4617dfe3b89f197eb17c112cce1a0d0c388320c9ff23d754f6150b035", "canonical_id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password", "child_ids": ["xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password", "parent_id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login", "path": "docs/guides/data-sources--api_crawler--properties--domains--simple_login--password.md", "provider_name": "api_crawler", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "simple_login", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/properties/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.simple_login.password for xcsh_api_crawler.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# domains.simple_login.password

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md)
- [Property reference](data-sources--api_crawler--reference.md)
- [domains](data-sources--api_crawler--properties--domains.md)
- [domains.simple_login](data-sources--api_crawler--properties--domains--simple_login.md)
- domains.simple_login.password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--api_crawler--properties--domains--simple_login--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--api_crawler--properties--domains--simple_login--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [domains.simple_login.password.blindfold_secret_info](data-sources--api_crawler--properties--domains--simple_login--password--blindfold_secret_info.md)
- [domains.simple_login.password.clear_secret_info](data-sources--api_crawler--properties--domains--simple_login--password--clear_secret_info.md)
- [domains.simple_login](data-sources--api_crawler--properties--domains--simple_login.md)
- [xcsh_api_crawler](../data-sources/api_crawler.md)
