---
page_title: "password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["password"], "body_bytes": 1956, "body_sha256": "sha256:ae1924e12802d41f940e397bd241cba3247f3b64bbd2961d08e1bebf9f0b1360", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:cminstance:properties:password:blindfold_secret_info", "xcsh-docs:resources:cminstance:properties:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:properties:password", "parent_id": "xcsh-docs:resources:cminstance:reference", "path": "documentation/resources/cminstance/properties/password/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2133031010321330-0110012311313220-1021213332102332-1032201111113230-0123321302231123-0130032231323311-3003311121000102-2321032213111120", "registry_path": "docs/guides/resources--cminstance--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cminstance:properties:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cminstance:properties:password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["password"], "schema_version": 1, "sections": [{"aliases": ["password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:cminstance:properties:password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cminstance:properties:password:blindfold_secret_info", "type": "requires"}], "schema_path": ["password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:cminstance:properties:password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-password--clear_secret_info--url", "enforcement": "provider-schema", "group": "password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cminstance:properties:password:clear_secret_info", "type": "requires"}], "schema_path": ["password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/properties/password/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
- password

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/): complete subsection reference.

## Next pages

- [password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/)
- [password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
