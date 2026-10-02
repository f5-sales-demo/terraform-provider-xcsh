---
page_title: "secondary.tsig_key_value"
subcategory: "DNS"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["secondary tsig key value"], "body_bytes": 2183, "body_sha256": "sha256:08461f0a664cfc899a28846ef729780d7164a52252ec385c69d9643bed1fba9f", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value", "parent_id": "xcsh-docs:resources:dns_zone:properties:secondary", "path": "documentation/resources/dns_zone/properties/secondary/tsig_key_value/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "secondary.tsig_key_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "secondary.tsig_key_value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["secondary", "tsig_key_value"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-secondary--tsig_key_value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "secondary.tsig_key_value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "type": "requires"}], "schema_path": ["secondary", "tsig_key_value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-secondary--tsig_key_value--clear_secret_info--url", "enforcement": "provider-schema", "group": "secondary.tsig_key_value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info", "type": "requires"}], "schema_path": ["secondary", "tsig_key_value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/secondary/tsig_key_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# secondary.tsig_key_value

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/)
- secondary.tsig_key_value

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
tsig_key_value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/): complete subsection reference.

## Next pages

- [secondary.tsig_key_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/)
- [secondary.tsig_key_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
