---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header"
subcategory: ""
description: "Header. Header matcher."
xcsh_docs: {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints header"], "body_bytes": 4588, "body_sha256": "sha256:0e3020a6aa0b07eb2ee1c94794fadf8ce0da88d909dac356a675cb5ede33fe49", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:all_header", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:header_and", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:header_none", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:header_or", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:not_present_header"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2301102211002031-2100310313001232-0330003002220230-2222000111212131-2032120112032131-2102322210021132-2101330302101311-1230122103100020", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header"], "schema_version": 1, "sections": [{"aliases": ["all header"], "anchor": "section", "description": "Can be used for the header operator choice where no values are needed.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:all_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header", "all_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["header and"], "anchor": "section", "description": "Request Body Matcher(s). A list of Header Matcher.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:header_and", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header", "header_and"], "syntax": "attribute", "type": "object"}, {"aliases": ["header none"], "anchor": "section", "description": "Configuration parameter for header none.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:header_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header", "header_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["header or"], "anchor": "section", "description": "Request Body Matcher(s). A list of Header Matcher.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:header_or", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header", "header_or"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--header--name", "description": "Name. Operator Name.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["not present header"], "anchor": "section", "description": "Can be used for the header operator choice where no values are needed.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header:not_present_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "header", "not_present_header"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Header. Header matcher.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header

<a id="section"></a>

Type: `"list"`. Computed.

Header. Header matcher.

## Direct properties

- [all_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/all_header/): complete subsection reference.

- [header_and](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/header_and/): complete subsection reference.

- [header_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/header_none/): complete subsection reference.

- [header_or](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/header_or/): complete subsection reference.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--header--name"></a>

### name property

Type: `"string"`. Computed.

Name. Operator Name.

- [not_present_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/not_present_header/): complete subsection reference.

## Next pages

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header.all_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/all_header/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header.header_and](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/header_and/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header.header_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/header_none/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header.header_or](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/header_or/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header.not_present_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/header/not_present_header/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
