---
page_title: "primary.dnssec_mode.enable"
subcategory: "DNS"
description: "primary.dnssec_mode.enable for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 894, "body_sha256": "sha256:7339072d2b6f89475fbacd79e9e580a86c434a317a7b1b4615d54f5880e07182", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "path": "docs/guides/resources--dns_zone--properties--primary--dnssec_mode--enable.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "dnssec_mode", "enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/dnssec_mode/enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.dnssec_mode.enable for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# primary.dnssec_mode.enable

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- [primary.dnssec_mode](resources--dns_zone--properties--primary--dnssec_mode.md)
- primary.dnssec_mode.enable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable. DNSSEC enable.

Upstream description:

DNSSEC enable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [primary.dnssec_mode](resources--dns_zone--properties--primary--dnssec_mode.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
