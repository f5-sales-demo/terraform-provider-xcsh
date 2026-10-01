---
page_title: "primary"
subcategory: "DNS"
description: "primary for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 2535, "body_sha256": "sha256:27b52ba78d21c274257f304e1758163ad931b1ecebe1764f82fe63eae8807774", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "xcsh-docs:data-sources:dns_zone:properties:primary:default_soa_parameters", "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group", "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary", "parent_id": "xcsh-docs:data-sources:dns_zone:reference", "path": "docs/guides/data-sources--dns_zone--properties--primary.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- primary

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: primary, secondary\] PrimaryDNSCreateSpecType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-soa_record_parameters_choice": "[\"default_soa_parameters\",\"soa_parameters\"]"
}
```

OneOf alternatives in this subsection:

- [primary](data-sources--dns_zone--properties--primary.md#section)
- [secondary](data-sources--dns_zone--properties--secondary.md#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-primary--allow_http_lb_managed_records"></a>

### allow_http_lb_managed_records property

Type: `"bool"`. Computed.

Option to allow user-created HTTP, TCP, and CDN load balancer related resource records to be
automatically managed in a protected RRset.

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

- [default_rr_set_group](data-sources--dns_zone--properties--primary--default_rr_set_group.md): complete subsection reference.

- [default_soa_parameters](data-sources--dns_zone--properties--primary--default_soa_parameters.md): complete subsection reference.

- [dnssec_mode](data-sources--dns_zone--properties--primary--dnssec_mode.md): complete subsection reference.

- [rr_set_group](data-sources--dns_zone--properties--primary--rr_set_group.md): complete subsection reference.

- [soa_parameters](data-sources--dns_zone--properties--primary--soa_parameters.md): complete subsection reference.

## Next pages

- [primary.default_rr_set_group](data-sources--dns_zone--properties--primary--default_rr_set_group.md)
- [primary.default_soa_parameters](data-sources--dns_zone--properties--primary--default_soa_parameters.md)
- [primary.dnssec_mode](data-sources--dns_zone--properties--primary--dnssec_mode.md)
- [primary.rr_set_group](data-sources--dns_zone--properties--primary--rr_set_group.md)
- [primary.soa_parameters](data-sources--dns_zone--properties--primary--soa_parameters.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
