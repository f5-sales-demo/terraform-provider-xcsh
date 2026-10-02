---
page_title: "domains.simple_login.password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["domains simple login password", "login", "login result", "sign in"], "body_bytes": 2130, "body_sha256": "sha256:82ec1a36cb7d21ae116afab7b5dd3a901a7614a025386191c0b9c011b2fd18a0", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password", "parent_id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login", "path": "documentation/data-sources/api_crawler/properties/domains/simple_login/password/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1102201101032321-2011133200223223-0230103103000213-3130332132100303-0321113030301111-3221031210122302-1221021320330310-1220220132312302", "registry_path": "docs/guides/data-sources--api_crawler--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "simple_login", "password"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "simple_login", "password", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:api_crawler:properties:domains:simple_login:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "simple_login", "password", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_crawler/properties/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.simple_login.password

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/)
- [domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/password/clear_secret_info/): complete subsection reference.

## Next pages

- [domains.simple_login.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/)
- [domains.simple_login.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/password/clear_secret_info/)
- [domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/properties/domains/simple_login/)
- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_crawler/)
