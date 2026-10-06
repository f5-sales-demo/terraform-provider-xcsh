---
page_title: "domains.simple_login.password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["domains simple login password", "login", "login result", "sign in"], "body_bytes": 1761, "body_sha256": "sha256:6ccf185f5f397b92aa5b739914ffead3cda92b5cf6d9c9efc37a32654f21d9a4", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password", "parent_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login", "path": "documentation/resources/api_crawler/properties/domains/simple_login/password/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2130123112313201-0122233002333311-0312121203002003-3030211023311213-1111011212323013-2032020232011131-1330210012221201-3212030030212111", "registry_path": "docs/guides/resources--api_crawler--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "domains.simple_login.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.simple_login.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "simple_login", "password"], "schema_version": 1, "sections": [{"aliases": ["domains simple login password blindfold secret info", "login", "login result", "sign in"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--simple_login--password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "domains.simple_login.password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:blindfold_secret_info", "type": "requires"}], "schema_path": ["domains", "simple_login", "password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["domains simple login password clear secret info", "login", "login result", "sign in"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--simple_login--password--clear_secret_info--url", "enforcement": "provider-schema", "group": "domains.simple_login.password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_crawler:properties:domains:simple_login:password:clear_secret_info", "type": "requires"}], "schema_path": ["domains", "simple_login", "password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/properties/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.simple_login.password

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/)
- [domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/clear_secret_info/): complete subsection reference.
