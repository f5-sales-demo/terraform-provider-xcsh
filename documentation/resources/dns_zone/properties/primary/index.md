---
page_title: "primary"
subcategory: "DNS"
description: "PrimaryDNSCreateSpecType."
xcsh_docs: {"aliases": ["primary"], "body_bytes": 3599, "body_sha256": "sha256:f6c793c12c104432d2702662d4fdb8aba197504e06f7cde1892a7b3d9bc6eee0", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group", "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary", "parent_id": "xcsh-docs:resources:dns_zone:reference", "path": "documentation/resources/dns_zone/properties/primary/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312", "registry_path": "docs/guides/resources--dns_zone--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary:ConflictingObjectAttributes:default_soa_parameters,soa_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary:ConflictingObjectAttributes:default_soa_parameters,soa_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary"], "schema_version": 1, "sections": [{"aliases": ["allow http lb managed records"], "anchor": "schema-primary--allow_http_lb_managed_records", "description": "Option to allow user-created HTTP, TCP, and CDN load balancer related resource records to be automatically managed in a protected RRset.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "allow_http_lb_managed_records"], "syntax": "attribute", "type": "bool"}, {"aliases": ["default rr set group"], "anchor": "section", "description": "Add and manage DNS resource record sets part of Default set group.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["primary", "default_rr_set_group"], "syntax": "block", "type": "object"}, {"aliases": ["default soa parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_soa_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_soa_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["dnssec mode"], "anchor": "section", "description": "DNSSEC Mode.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.dnssec_mode:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.dnssec_mode:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "type": "conflicts"}], "schema_path": ["primary", "dnssec_mode"], "syntax": "block", "type": "object"}, {"aliases": ["rr set group"], "anchor": "section", "description": "Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed by F5.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["primary", "rr_set_group"], "syntax": "block", "type": "object"}, {"aliases": ["soa parameters"], "anchor": "section", "description": "Configuration parameter for soa parameters.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--soa_parameters--refresh", "enforcement": "provider-schema", "group": "primary.soa_parameters:RequiredObjectAttributes:refresh,retry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "type": "requires"}, {"anchor": "schema-primary--soa_parameters--retry", "enforcement": "provider-schema", "group": "primary.soa_parameters:RequiredObjectAttributes:refresh,retry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "type": "requires"}], "schema_path": ["primary", "soa_parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "PrimaryDNSCreateSpecType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
