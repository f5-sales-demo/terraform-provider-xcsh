---
page_title: "secondary.tsig_key_value"
subcategory: "DNS"
description: "secondary.tsig_key_value for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1446, "body_sha256": "sha256:d074f1428955faf639a40baa748684d16f342bd63a9c1bdcbba8c23f08ec93ed", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:blindfold_secret_info", "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "path": "docs/guides/data-sources--dns_zone--properties--secondary--tsig_key_value.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["secondary", "tsig_key_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/secondary/tsig_key_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "secondary.tsig_key_value for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# secondary.tsig_key_value

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [secondary](data-sources--dns_zone--properties--secondary.md)
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

- [blindfold_secret_info](data-sources--dns_zone--properties--secondary--tsig_key_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--dns_zone--properties--secondary--tsig_key_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [secondary.tsig_key_value.blindfold_secret_info](data-sources--dns_zone--properties--secondary--tsig_key_value--blindfold_secret_info.md)
- [secondary.tsig_key_value.clear_secret_info](data-sources--dns_zone--properties--secondary--tsig_key_value--clear_secret_info.md)
- [secondary](data-sources--dns_zone--properties--secondary.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
