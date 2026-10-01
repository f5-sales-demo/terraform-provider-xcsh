---
page_title: "primary.default_rr_set_group"
subcategory: "DNS"
description: "primary.default_rr_set_group for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 7767, "body_sha256": "sha256:b88d95fdf7eac7714139967f4e4284bcbb4c1c4da4b47bff7f8dd83c6e37ad0d", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:a_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:aaaa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:afsdb_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:alias_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:caa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cert_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cname_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ds_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:eui48_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:eui64_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:lb_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:loc_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:mx_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:naptr_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ns_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ptr_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:srv_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:sshfp_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:tlsa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:txt_record"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "path": "docs/guides/data-sources--dns_zone--properties--primary--default_rr_set_group.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "default_rr_set_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- primary.default_rr_set_group

<a id="section"></a>

Type: `"list"`. Computed.

Add and manage DNS resource record sets part of Default set group.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

## Direct properties

- [a_record](data-sources--dns_zone--properties--primary--default_rr_set_group--a_record.md): complete subsection reference.

- [aaaa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--aaaa_record.md): complete subsection reference.

- [afsdb_record](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record.md): complete subsection reference.

- [alias_record](data-sources--dns_zone--properties--primary--default_rr_set_group--alias_record.md): complete subsection reference.

- [caa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record.md): complete subsection reference.

- [cds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record.md): complete subsection reference.

- [cert_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record.md): complete subsection reference.

- [cname_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cname_record.md): complete subsection reference.

<a id="schema-primary--default_rr_set_group--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record.md): complete subsection reference.

- [eui48_record](data-sources--dns_zone--properties--primary--default_rr_set_group--eui48_record.md): complete subsection reference.

- [eui64_record](data-sources--dns_zone--properties--primary--default_rr_set_group--eui64_record.md): complete subsection reference.

- [lb_record](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record.md): complete subsection reference.

- [loc_record](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record.md): complete subsection reference.

- [mx_record](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record.md): complete subsection reference.

- [naptr_record](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record.md): complete subsection reference.

- [ns_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ns_record.md): complete subsection reference.

- [ptr_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ptr_record.md): complete subsection reference.

- [srv_record](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record.md): complete subsection reference.

- [sshfp_record](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record.md): complete subsection reference.

- [tlsa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record.md): complete subsection reference.

<a id="schema-primary--default_rr_set_group--ttl"></a>

### ttl property

Type: `"number"`. Computed.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](data-sources--dns_zone--properties--primary--default_rr_set_group--txt_record.md): complete subsection reference.

## Next pages

- [primary.default_rr_set_group.a_record](data-sources--dns_zone--properties--primary--default_rr_set_group--a_record.md)
- [primary.default_rr_set_group.aaaa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--aaaa_record.md)
- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record.md)
- [primary.default_rr_set_group.alias_record](data-sources--dns_zone--properties--primary--default_rr_set_group--alias_record.md)
- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record.md)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record.md)
- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record.md)
- [primary.default_rr_set_group.cname_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cname_record.md)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record.md)
- [primary.default_rr_set_group.eui48_record](data-sources--dns_zone--properties--primary--default_rr_set_group--eui48_record.md)
- [primary.default_rr_set_group.eui64_record](data-sources--dns_zone--properties--primary--default_rr_set_group--eui64_record.md)
- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record.md)
- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record.md)
- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record.md)
- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record.md)
- [primary.default_rr_set_group.ns_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ns_record.md)
- [primary.default_rr_set_group.ptr_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ptr_record.md)
- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record.md)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record.md)
- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record.md)
- [primary.default_rr_set_group.txt_record](data-sources--dns_zone--properties--primary--default_rr_set_group--txt_record.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
