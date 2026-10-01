---
page_title: "secondary.tsig_key_value"
subcategory: "DNS"
description: "secondary.tsig_key_value for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1899, "body_sha256": "sha256:85de036a06d4f23ee9acedb1e9440ad44221152e054424fc47cd6f23d167c0be", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "path": "documentation/data-sources/dns_zone/properties/secondary/tsig_key_value/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["secondary", "tsig_key_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/secondary/tsig_key_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "secondary.tsig_key_value for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
