---
page_title: "primary.rr_set_group.rr_set"
subcategory: "DNS"
description: "primary.rr_set_group.rr_set for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 24774, "body_sha256": "sha256:67608308861b88b0ee013f5eb85ce06cb2469f39b0265e8e66b4ce3871723d14", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:a_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:aaaa_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:alias_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:caa_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cert_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cname_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:ds_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:eui48_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:eui64_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:lb_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:loc_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:mx_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:ns_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:ptr_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:srv_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:txt_record"], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group", "path": "docs/guides/resources--dns_zone--properties--primary--rr_set_group--rr_set.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.rr_set_group.rr_set for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- [primary.rr_set_group](resources--dns_zone--properties--primary--rr_set_group.md)
- primary.rr_set_group.rr_set

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Resource Record Sets. Collection of DNS resource record sets.

Upstream description:

Collection of DNS resource record sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ttl"),
  validators.ConflictingListObjectAttributes("a_record",
    "aaaa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("tlsa_record",
    "txt_record")}
```

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

Terraform syntax:

```terraform
rr_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [a_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--a_record.md): complete subsection reference.

- [aaaa_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--aaaa_record.md): complete subsection reference.

- [afsdb_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record.md): complete subsection reference.

- [alias_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--alias_record.md): complete subsection reference.

- [caa_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record.md): complete subsection reference.

- [cds_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record.md): complete subsection reference.

- [cert_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record.md): complete subsection reference.

- [cname_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--cname_record.md): complete subsection reference.

<a id="schema-primary--rr_set_group--rr_set--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Comment. Human-readable description text

- [ds_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record.md): complete subsection reference.

- [eui48_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--eui48_record.md): complete subsection reference.

- [eui64_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--eui64_record.md): complete subsection reference.

- [lb_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record.md): complete subsection reference.

- [loc_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record.md): complete subsection reference.

- [mx_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record.md): complete subsection reference.

- [naptr_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record.md): complete subsection reference.

- [ns_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--ns_record.md): complete subsection reference.

- [ptr_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--ptr_record.md): complete subsection reference.

- [srv_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record.md): complete subsection reference.

- [sshfp_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record.md): complete subsection reference.

- [tlsa_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record.md): complete subsection reference.

<a id="schema-primary--rr_set_group--rr_set--ttl"></a>

### ttl property

Type: `"number"`. Optional.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

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

- [txt_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--txt_record.md): complete subsection reference.

## Next pages

- [primary.rr_set_group.rr_set.a_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--a_record.md)
- [primary.rr_set_group.rr_set.aaaa_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--aaaa_record.md)
- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record.md)
- [primary.rr_set_group.rr_set.alias_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--alias_record.md)
- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record.md)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record.md)
- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record.md)
- [primary.rr_set_group.rr_set.cname_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--cname_record.md)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record.md)
- [primary.rr_set_group.rr_set.eui48_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--eui48_record.md)
- [primary.rr_set_group.rr_set.eui64_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--eui64_record.md)
- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record.md)
- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record.md)
- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record.md)
- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record.md)
- [primary.rr_set_group.rr_set.ns_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--ns_record.md)
- [primary.rr_set_group.rr_set.ptr_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--ptr_record.md)
- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record.md)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record.md)
- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record.md)
- [primary.rr_set_group.rr_set.txt_record](resources--dns_zone--properties--primary--rr_set_group--rr_set--txt_record.md)
- [primary.rr_set_group](resources--dns_zone--properties--primary--rr_set_group.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
