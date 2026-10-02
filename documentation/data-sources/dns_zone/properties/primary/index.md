---
page_title: "primary"
subcategory: "DNS"
description: "PrimaryDNSCreateSpecType."
xcsh_docs: {"aliases": ["primary"], "body_bytes": 3345, "body_sha256": "sha256:6e43a99de5a3bae612e82708b472eb95b7fbc0e38cf1eb7e796ddd497ee8a9bf", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "xcsh-docs:data-sources:dns_zone:properties:primary:default_soa_parameters", "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode", "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group", "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary", "parent_id": "xcsh-docs:data-sources:dns_zone:reference", "path": "documentation/data-sources/dns_zone/properties/primary/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary"], "schema_version": 1, "sections": [{"aliases": ["allow http lb managed records"], "anchor": "schema-primary--allow_http_lb_managed_records", "description": "Option to allow user-created HTTP, TCP, and CDN load balancer related resource records to be automatically managed in a protected RRset.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "allow_http_lb_managed_records"], "syntax": "attribute", "type": "bool"}, {"aliases": ["default rr set group"], "anchor": "section", "description": "Add and manage DNS resource record sets part of Default set group.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["primary", "default_rr_set_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["default soa parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_soa_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["dnssec mode"], "anchor": "section", "description": "DNSSEC Mode.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["primary", "dnssec_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["rr set group"], "anchor": "section", "description": "Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed by F5.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["primary", "rr_set_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["soa parameters"], "anchor": "section", "description": "Configuration parameter for soa parameters.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["primary", "soa_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "PrimaryDNSCreateSpecType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
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

- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/#section)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/#section)

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

- [default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/): complete subsection reference.

- [default_soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_soa_parameters/): complete subsection reference.

- [dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/): complete subsection reference.

- [rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/): complete subsection reference.

- [soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/): complete subsection reference.

## Next pages

- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_soa_parameters/)
- [primary.dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/)
- [primary.soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
