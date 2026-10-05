---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_threats."
xcsh_docs: {"aliases": ["waf threats"], "body_bytes": 3117, "body_sha256": "sha256:3c40e848203f9b6e48b6b1f8c1e8dead7cb4ceb4fe0b5df1e74421d4999d5117", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_threats:properties:cve_ids"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:reference", "parent_id": "xcsh-docs:data-sources:waf_threats:fundamentals", "path": "documentation/data-sources/waf_threats/properties/index.md", "product": "distributed-cloud", "provider_name": "waf_threats", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1221102221200320-3320211232000012-1012023131032222-0213323223221202-3331003100012130-1302101323323020-0023211201200003-3020110212312032", "registry_path": "docs/guides/data-sources--waf_threats--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["cursor"], "anchor": "schema-cursor", "description": "Opaque pagination cursor returned from previous response.", "document_id": "xcsh-docs:data-sources:waf_threats:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cursor"], "syntax": "attribute", "type": "string"}, {"aliases": ["cve ids"], "anchor": "section", "description": "CVE ID List. A list of CVE IDs.", "document_id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cve_ids"], "syntax": "attribute", "type": "object"}, {"aliases": ["next cursor"], "anchor": "schema-next_cursor", "description": "Next Cursor. Opaque cursor for fetching next page.", "document_id": "xcsh-docs:data-sources:waf_threats:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["next_cursor"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary tag"], "anchor": "schema-primary_tag", "description": "Exclusive with Primary tag to filter threats. A primary tag is a high level categorization of a threat, such as an associated threat actor, malware family or CVE.", "document_id": "xcsh-docs:data-sources:waf_threats:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["report fields"], "anchor": "schema-report_fields", "description": "Optional list of fields to include in threat representation.", "document_id": "xcsh-docs:data-sources:waf_threats:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["report_fields"], "syntax": "attribute", "type": "list"}, {"aliases": ["threats"], "anchor": "schema-threats", "description": "Threats. A list of threats that match the query.", "document_id": "xcsh-docs:data-sources:waf_threats:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threats"], "syntax": "attribute", "type": "list"}, {"aliases": ["waf sec event id"], "anchor": "schema-waf_sec_event_id", "description": "Exclusive with WAF Security Event ID to GET associated threats information.", "document_id": "xcsh-docs:data-sources:waf_threats:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_sec_event_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_waf_threats.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_waf_threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/)
- Property reference

## Direct properties

<a id="schema-cursor"></a>

### cursor property

Type: `"string"`. Optional.

Opaque pagination cursor returned from previous response.

- [cve_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/cve_ids/): complete subsection reference.

<a id="schema-next_cursor"></a>

### next_cursor property

Type: `"string"`. Computed.

Next Cursor. Opaque cursor for fetching next page.

<a id="schema-primary_tag"></a>

### primary_tag property

Type: `"string"`. Optional.

Exclusive with \[cve\_ids waf\_sec\_event\_id\] Primary tag to filter threats. A primary tag is a
high level categorization of a threat, such as an associated threat actor, malware family or CVE.

<a id="schema-report_fields"></a>

### report_fields property

Type: `["list", "string"]`. Optional.

Optional list of fields to include in threat representation.

<a id="schema-threats"></a>

### threats property

Type: `["list", "string"]`. Computed.

Threats. A list of threats that match the query.

<a id="schema-waf_sec_event_id"></a>

### waf_sec_event_id property

Type: `"string"`. Optional.

Exclusive with \[cve\_ids primary\_tag\] WAF Security Event ID to GET associated threats
information.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cursor` | [cursor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/#schema-cursor) |
| `cve_ids` | [cve_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/cve_ids/#section) |
| `cve_ids.ids` | [cve_ids.ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/cve_ids/#schema-cve_ids--ids) |
| `next_cursor` | [next_cursor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/#schema-next_cursor) |
| `primary_tag` | [primary_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/#schema-primary_tag) |
| `report_fields` | [report_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/#schema-report_fields) |
| `threats` | [threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/#schema-threats) |
| `waf_sec_event_id` | [waf_sec_event_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/#schema-waf_sec_event_id) |

## Next pages

- [cve_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/cve_ids/)
- [xcsh_waf_threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/)
