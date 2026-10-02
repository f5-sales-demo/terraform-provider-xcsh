---
page_title: "code_base_integration.bitbucket.passwd"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["code base integration bitbucket passwd"], "body_bytes": 2667, "body_sha256": "sha256:eae3515726971af376a322f655802a842d7b60bfc873b281618d616544b0461a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:blindfold_secret_info", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket", "path": "documentation/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303", "registry_path": "docs/guides/resources--code_base_integration--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket.passwd:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket.passwd:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "bitbucket", "passwd"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-code_base_integration--bitbucket--passwd--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket.passwd.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:blindfold_secret_info", "type": "requires"}], "schema_path": ["code_base_integration", "bitbucket", "passwd", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-code_base_integration--bitbucket--passwd--clear_secret_info--url", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket.passwd.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket:passwd:clear_secret_info", "type": "requires"}], "schema_path": ["code_base_integration", "bitbucket", "passwd", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.bitbucket.passwd

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- [code_base_integration.bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/)
- code_base_integration.bitbucket.passwd

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/clear_secret_info/): complete subsection reference.

## Next pages

- [code_base_integration.bitbucket.passwd.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/blindfold_secret_info/)
- [code_base_integration.bitbucket.passwd.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/passwd/clear_secret_info/)
- [code_base_integration.bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
