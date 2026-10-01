---
page_title: "primary.rr_set_group"
subcategory: "DNS"
description: "primary.rr_set_group for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1922, "body_sha256": "sha256:3db029531327d8b4165b62d8ca825539b65177bb3de6b0139f2c5b4c6ce8519d", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:metadata", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set"], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "docs/guides/resources--dns_zone--properties--primary--rr_set_group.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "rr_set_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.rr_set_group for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- primary.rr_set_group

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rr_set_group {
  # Configure direct properties listed below.
}
```

## Direct properties

- [metadata](resources--dns_zone--properties--primary--rr_set_group--metadata.md): complete subsection reference.

- [rr_set](resources--dns_zone--properties--primary--rr_set_group--rr_set.md): complete subsection reference.

## Next pages

- [primary.rr_set_group.metadata](resources--dns_zone--properties--primary--rr_set_group--metadata.md)
- [primary.rr_set_group.rr_set](resources--dns_zone--properties--primary--rr_set_group--rr_set.md)
- [primary](resources--dns_zone--properties--primary.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
