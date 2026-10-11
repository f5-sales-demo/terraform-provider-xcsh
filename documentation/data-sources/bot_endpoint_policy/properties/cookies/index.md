---
page_title: "cookies"
subcategory: ""
description: "The cookie names and corresponding descriptions for this policy."
xcsh_docs: {"aliases": ["cookies"], "body_bytes": 792, "body_sha256": "sha256:8113a05fef0a6b8e4c1e1b5efacb9b03f679cd8cf3abda54df0e9e702b90ea53", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:cookies", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:reference", "path": "documentation/data-sources/bot_endpoint_policy/properties/cookies/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3000112311223122-0232231223012220-0123212120032310-3331011121112331-0131230122211201-1313101013312003-3220313323103211-0130230023133300", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cookies"], "schema_version": 1, "sections": [{"aliases": ["cookies description spec"], "anchor": "schema-cookies--description_spec", "description": "Human‐readable explanation of what the cookie does.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:cookies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookies", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["cookies name"], "anchor": "schema-cookies--name", "description": "Name. The name of the cookie.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:cookies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cookies", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/cookies/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "The cookie names and corresponding descriptions for this policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookies

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- cookies

<a id="section"></a>

Type: `"list"`. Computed.

The cookie names and corresponding descriptions for this policy.

## Direct properties

<a id="schema-cookies--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Human‐readable explanation of what the cookie does.

<a id="schema-cookies--name"></a>

### name property

Type: `"string"`. Computed.

Name. The name of the cookie.
