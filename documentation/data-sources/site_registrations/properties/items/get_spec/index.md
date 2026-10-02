---
page_title: "items.get_spec"
subcategory: ""
description: "GET Registration. GET registration specification."
xcsh_docs: {"aliases": ["items get spec"], "body_bytes": 1665, "body_sha256": "sha256:e62b54d128661be30c5f8682eb63c917e62a097498480c338d7a1cdd3cbe5b8a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:passport"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3333130102202200-3111230132200003-3031121301303312-2330122102310202-2323111021200031-2221112021221301-1032220031102133-0110033201202331", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec"], "schema_version": 1, "sections": [{"aliases": ["infra"], "anchor": "section", "description": "InfraMetadata stores information about instance infrastructure.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra"], "syntax": "attribute", "type": "object"}, {"aliases": ["passport"], "anchor": "section", "description": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:passport", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "passport"], "syntax": "attribute", "type": "object"}, {"aliases": ["token"], "anchor": "schema-items--get_spec--token", "description": "Token is used for machine and tenant identification.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "token"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "GET Registration. GET registration specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- items.get_spec

<a id="section"></a>

Type: `"single"`. Computed.

GET Registration. GET registration specification.

## Direct properties

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/): complete subsection reference.

- [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/passport/): complete subsection reference.

<a id="schema-items--get_spec--token"></a>

### token property

Type: `"string"`. Computed.

Token is used for machine and tenant identification.

## Next pages

- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/passport/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
