---
page_title: "primary.dnssec_mode"
subcategory: "DNS"
description: "primary.dnssec_mode for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1832, "body_sha256": "sha256:8ac03f1a8690907677facd4c9081c525a30d72aaf942ad405b1afcf1de6b2cd2", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable"], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "documentation/resources/dns_zone/properties/primary/dnssec_mode/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["primary", "dnssec_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/dnssec_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.dnssec_mode for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# primary.dnssec_mode

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
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

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/enable/): complete subsection reference.

## Next pages

- [primary.dnssec_mode.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/disable_spec/)
- [primary.dnssec_mode.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/enable/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
