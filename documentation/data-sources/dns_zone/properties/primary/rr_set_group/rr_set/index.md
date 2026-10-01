---
page_title: "primary.rr_set_group.rr_set"
subcategory: "DNS"
description: "primary.rr_set_group.rr_set for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 10239, "body_sha256": "sha256:b6119576debb64e452c805c438db0564a7127c8961f6e1b26d41be688892a99c", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:a_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:aaaa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:alias_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:caa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:cert_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:cname_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:ds_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:eui48_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:eui64_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:lb_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:loc_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:mx_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:ns_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:ptr_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:srv_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:tlsa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:txt_record"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group", "path": "documentation/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.rr_set_group.rr_set for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/)
- primary.rr_set_group.rr_set

<a id="section"></a>

Type: `"list"`. Computed.

Resource Record Sets. Collection of DNS resource record sets.

Upstream description:

Collection of DNS resource record sets.

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

- [a_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/a_record/): complete subsection reference.

- [aaaa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/aaaa_record/): complete subsection reference.

- [afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/): complete subsection reference.

- [alias_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/alias_record/): complete subsection reference.

- [caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/): complete subsection reference.

- [cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/): complete subsection reference.

- [cert_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/): complete subsection reference.

- [cname_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cname_record/): complete subsection reference.

<a id="schema-primary--rr_set_group--rr_set--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/): complete subsection reference.

- [eui48_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui48_record/): complete subsection reference.

- [eui64_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/): complete subsection reference.

- [lb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/): complete subsection reference.

- [loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/): complete subsection reference.

- [mx_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/): complete subsection reference.

- [naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/): complete subsection reference.

- [ns_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ns_record/): complete subsection reference.

- [ptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ptr_record/): complete subsection reference.

- [srv_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/): complete subsection reference.

- [sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/): complete subsection reference.

- [tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/): complete subsection reference.

<a id="schema-primary--rr_set_group--rr_set--ttl"></a>

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

- [txt_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/txt_record/): complete subsection reference.

## Next pages

- [primary.rr_set_group.rr_set.a_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/a_record/)
- [primary.rr_set_group.rr_set.aaaa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/aaaa_record/)
- [primary.rr_set_group.rr_set.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/)
- [primary.rr_set_group.rr_set.alias_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/alias_record/)
- [primary.rr_set_group.rr_set.caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/)
- [primary.rr_set_group.rr_set.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/)
- [primary.rr_set_group.rr_set.cert_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/)
- [primary.rr_set_group.rr_set.cname_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cname_record/)
- [primary.rr_set_group.rr_set.ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/)
- [primary.rr_set_group.rr_set.eui48_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui48_record/)
- [primary.rr_set_group.rr_set.eui64_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/)
- [primary.rr_set_group.rr_set.lb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/)
- [primary.rr_set_group.rr_set.loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/)
- [primary.rr_set_group.rr_set.mx_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/)
- [primary.rr_set_group.rr_set.naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/)
- [primary.rr_set_group.rr_set.ns_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ns_record/)
- [primary.rr_set_group.rr_set.ptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ptr_record/)
- [primary.rr_set_group.rr_set.srv_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/)
- [primary.rr_set_group.rr_set.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/)
- [primary.rr_set_group.rr_set.tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/)
- [primary.rr_set_group.rr_set.txt_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/txt_record/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
