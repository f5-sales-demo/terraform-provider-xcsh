---
page_title: "ips"
subcategory: ""
description: "List of IP addresses along with log count."
xcsh_docs: {"aliases": ["ips"], "body_bytes": 745, "body_sha256": "sha256:a378815652c2208e2baccb258e2195d1d063743d4ffbcb2dee68ef07bcf818b1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "path": "documentation/data-sources/infraprotect_mitigation_ips/properties/ips/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1333112312212100-0102213032232210-3210130003030011-0312201101131201-2202102211132312-0112112121120222-3131030131230322-0113013031121110", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ips"], "schema_version": 1, "sections": [{"aliases": ["ips ip"], "anchor": "schema-ips--ip", "description": "IP. Mitigation source IP.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["ips log count"], "anchor": "schema-ips--log_count", "description": "Number of times the IP appears in the log.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "log_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/properties/ips/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of IP addresses along with log count.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ips

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/)
- ips

<a id="section"></a>

Type: `"list"`. Computed.

List of IP addresses along with log count.

## Direct properties

<a id="schema-ips--ip"></a>

### ip property

Type: `"string"`. Computed.

IP. Mitigation source IP.

<a id="schema-ips--log_count"></a>

### log_count property

Type: `"string"`. Computed.

Number of times the IP appears in the log.
