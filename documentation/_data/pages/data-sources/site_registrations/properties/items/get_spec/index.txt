---
page_title: "items.get_spec"
subcategory: ""
description: "GET Registration. GET registration specification."
xcsh_docs: {"aliases": ["items get spec"], "body_bytes": 1114, "body_sha256": "sha256:fb75d2a0529c21481f34e88b5379bddd6eeec765b10f8dde495be0869a77e199", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:passport"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3333130102202200-3111230132200003-3031121301303312-2330122102310202-2323111021200031-2221112021221301-1032220031102133-0110033201202331", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra"], "anchor": "section", "description": "InfraMetadata stores information about instance infrastructure.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra"], "syntax": "attribute", "type": "object"}, {"aliases": ["items get spec passport"], "anchor": "section", "description": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "passport"], "syntax": "attribute", "type": "object"}, {"aliases": ["items get spec token"], "anchor": "schema-items--get_spec--token", "description": "Token is used for machine and tenant identification.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "token"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "GET Registration. GET registration specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
