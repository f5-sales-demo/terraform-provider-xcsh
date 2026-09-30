---
page_title: "primary.default_soa_parameters"
subcategory: "DNS"
description: "primary.default_soa_parameters for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 883, "body_sha256": "sha256:271b8214230707e6e65248189f6780a84f7adee4f6d4e487d3a11ac991f291ab", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "docs/guides/resources--dns_zone--properties--primary--default_soa_parameters.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "default_soa_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_soa_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_soa_parameters for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# primary.default_soa_parameters

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- primary.default_soa_parameters

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default soa parameters.

Upstream description:

This can be used for messages where no values are needed.

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
default_soa_parameters = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [primary](resources--dns_zone--properties--primary.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
