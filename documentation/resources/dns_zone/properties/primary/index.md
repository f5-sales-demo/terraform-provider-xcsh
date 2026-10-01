---
page_title: "primary"
subcategory: "DNS"
description: "primary for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 3599, "body_sha256": "sha256:f6c793c12c104432d2702662d4fdb8aba197504e06f7cde1892a7b3d9bc6eee0", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group", "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters"], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary", "parent_id": "xcsh-docs:resources:dns_zone:reference", "path": "documentation/resources/dns_zone/properties/primary/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["primary"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- primary

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: primary, secondary\] PrimaryDNSCreateSpecType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_soa_parameters",
    "soa_parameters")}
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
  "x-ves-oneof-field-soa_record_parameters_choice": "[\"default_soa_parameters\",\"soa_parameters\"]"
}
```

OneOf alternatives in this subsection:

- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/#section)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/secondary/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
primary {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--allow_http_lb_managed_records"></a>

### allow_http_lb_managed_records property

Type: `"bool"`. Optional.

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

- [default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/): complete subsection reference.

- [default_soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_soa_parameters/): complete subsection reference.

- [dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/): complete subsection reference.

- [rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/): complete subsection reference.

- [soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/soa_parameters/): complete subsection reference.

## Next pages

- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_soa_parameters/)
- [primary.dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/soa_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
