---
page_title: "ips"
subcategory: ""
description: "List of IP addresses along with log count."
xcsh_docs: {"aliases": ["ips"], "body_bytes": 1030, "body_sha256": "sha256:e5f2f8712dd3952d8d7145be4765924cfbb4cd9e65dbb24fa18418711f3ab751", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "path": "documentation/data-sources/infraprotect_mitigation_ips/properties/ips/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1333112312212100-0102213032232210-3210130003030011-0312201101131201-2202102211132312-0112112121120222-3131030131230322-0113013031121110", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ips"], "schema_version": 1, "sections": [{"aliases": ["ip"], "anchor": "schema-ips--ip", "description": "IP. Mitigation source IP.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["log count"], "anchor": "schema-ips--log_count", "description": "Number of times the IP appears in the log.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "log_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/properties/ips/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IP addresses along with log count.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/)
- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
