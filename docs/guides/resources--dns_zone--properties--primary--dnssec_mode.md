---
page_title: "primary.dnssec_mode"
subcategory: "DNS"
description: "primary.dnssec_mode for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1478, "body_sha256": "sha256:d2753df6f4889c45ad29deee618c9e4ef94ef9cc675431f5c9c63044965d13c1", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable"], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "docs/guides/resources--dns_zone--properties--primary--dnssec_mode.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "dnssec_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/dnssec_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.dnssec_mode for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.dnssec_mode

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- primary.dnssec_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNSSEC Mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
dnssec_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](resources--dns_zone--properties--primary--dnssec_mode--disable_spec.md): complete subsection reference.

- [enable](resources--dns_zone--properties--primary--dnssec_mode--enable.md): complete subsection reference.

## Next pages

- [primary.dnssec_mode.disable_spec](resources--dns_zone--properties--primary--dnssec_mode--disable_spec.md)
- [primary.dnssec_mode.enable](resources--dns_zone--properties--primary--dnssec_mode--enable.md)
- [primary](resources--dns_zone--properties--primary.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
