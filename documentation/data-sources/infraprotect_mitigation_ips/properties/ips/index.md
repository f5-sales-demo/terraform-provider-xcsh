---
page_title: "ips"
subcategory: ""
description: "List of IP addresses along with log count."
xcsh_docs: {"aliases": ["ips"], "body_bytes": 1030, "body_sha256": "sha256:e5f2f8712dd3952d8d7145be4765924cfbb4cd9e65dbb24fa18418711f3ab751", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "path": "documentation/data-sources/infraprotect_mitigation_ips/properties/ips/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1333112312212100-0102213032232210-3210130003030011-0312201101131201-2202102211132312-0112112121120222-3131030131230322-0113013031121110", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ips"], "schema_version": 1, "sections": [{"aliases": ["ips ip"], "anchor": "schema-ips--ip", "description": "IP. Mitigation source IP.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["ips log count"], "anchor": "schema-ips--log_count", "description": "Number of times the IP appears in the log.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "log_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/properties/ips/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of IP addresses along with log count.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
