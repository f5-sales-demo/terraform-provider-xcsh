---
page_title: "secondary.tsig_key_value"
subcategory: "DNS"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["secondary tsig key value"], "body_bytes": 1899, "body_sha256": "sha256:85de036a06d4f23ee9acedb1e9440ad44221152e054424fc47cd6f23d167c0be", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "path": "documentation/data-sources/dns_zone/properties/secondary/tsig_key_value/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["secondary", "tsig_key_value"], "schema_version": 1, "sections": [{"aliases": ["secondary tsig key value blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["secondary", "tsig_key_value", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["secondary tsig key value clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["secondary", "tsig_key_value", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/secondary/tsig_key_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# secondary.tsig_key_value

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/)
- secondary.tsig_key_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/): complete subsection reference.

## Next pages

- [secondary.tsig_key_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/)
- [secondary.tsig_key_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
