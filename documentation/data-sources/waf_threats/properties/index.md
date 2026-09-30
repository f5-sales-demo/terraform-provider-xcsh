---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_threats."
xcsh_docs: {"aliases": [], "body_bytes": 3018, "body_sha256": "sha256:4f37e54c3a020706c4a819ba0fff0d848735f53dbae0e64d78649b0aef474ff1", "child_ids": ["xcsh-docs:data-sources:waf_threats:properties:cve_ids"], "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:reference", "parent_id": "xcsh-docs:data-sources:waf_threats:fundamentals", "path": "documentation/data-sources/waf_threats/properties/index.md", "provider_name": "waf_threats", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_waf_threats.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
