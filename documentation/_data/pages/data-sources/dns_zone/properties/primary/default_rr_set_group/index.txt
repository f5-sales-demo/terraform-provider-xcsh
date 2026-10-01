---
page_title: "primary.default_rr_set_group"
subcategory: "DNS"
description: "primary.default_rr_set_group for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 10082, "body_sha256": "sha256:5e449621d23328ec182c39ae95ad3ee8ed4cbd223f3fbea51592954dbc7a3409", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:a_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:aaaa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:afsdb_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:alias_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:caa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cert_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cname_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ds_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:eui48_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:eui64_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:lb_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:loc_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:mx_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:naptr_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ns_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ptr_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:srv_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:sshfp_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:tlsa_record", "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:txt_record"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "path": "documentation/data-sources/dns_zone/properties/primary/default_rr_set_group/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["primary", "default_rr_set_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
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

- [a_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/a_record/): complete subsection reference.

- [aaaa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/aaaa_record/): complete subsection reference.

- [afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/): complete subsection reference.

- [alias_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/alias_record/): complete subsection reference.

- [caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/): complete subsection reference.

- [cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/): complete subsection reference.

- [cert_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/): complete subsection reference.

- [cname_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cname_record/): complete subsection reference.

<a id="schema-primary--default_rr_set_group--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/): complete subsection reference.

- [eui48_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui48_record/): complete subsection reference.

- [eui64_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui64_record/): complete subsection reference.

- [lb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/): complete subsection reference.

- [loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/): complete subsection reference.

- [mx_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/): complete subsection reference.

- [naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/): complete subsection reference.

- [ns_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ns_record/): complete subsection reference.

- [ptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ptr_record/): complete subsection reference.

- [srv_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/): complete subsection reference.

- [sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/): complete subsection reference.

- [tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/): complete subsection reference.

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

- [txt_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/): complete subsection reference.

## Next pages

- [primary.default_rr_set_group.a_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/a_record/)
- [primary.default_rr_set_group.aaaa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/aaaa_record/)
- [primary.default_rr_set_group.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/)
- [primary.default_rr_set_group.alias_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/alias_record/)
- [primary.default_rr_set_group.caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/)
- [primary.default_rr_set_group.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/)
- [primary.default_rr_set_group.cert_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/)
- [primary.default_rr_set_group.cname_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cname_record/)
- [primary.default_rr_set_group.ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/)
- [primary.default_rr_set_group.eui48_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui48_record/)
- [primary.default_rr_set_group.eui64_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui64_record/)
- [primary.default_rr_set_group.lb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/)
- [primary.default_rr_set_group.loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/)
- [primary.default_rr_set_group.mx_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/)
- [primary.default_rr_set_group.naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/)
- [primary.default_rr_set_group.ns_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ns_record/)
- [primary.default_rr_set_group.ptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ptr_record/)
- [primary.default_rr_set_group.srv_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/)
- [primary.default_rr_set_group.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/)
- [primary.default_rr_set_group.tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/)
- [primary.default_rr_set_group.txt_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
