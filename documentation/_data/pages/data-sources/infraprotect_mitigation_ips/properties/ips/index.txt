---
page_title: "ips"
subcategory: ""
description: "List of IP addresses along with log count."
xcsh_docs: {"aliases": ["ips"], "body_bytes": 1030, "body_sha256": "sha256:e5f2f8712dd3952d8d7145be4765924cfbb4cd9e65dbb24fa18418711f3ab751", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "path": "documentation/data-sources/infraprotect_mitigation_ips/properties/ips/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1333112312212100-0102213032232210-3210130003030011-0312201101131201-2202102211132312-0112112121120222-3131030131230322-0113013031121110", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ips"], "schema_version": 1, "sections": [{"aliases": ["ip"], "anchor": "schema-ips--ip", "description": "IP. Mitigation source IP.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["log count"], "anchor": "schema-ips--log_count", "description": "Number of times the IP appears in the log.", "document_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:properties:ips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ips", "log_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/properties/ips/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of IP addresses along with log count.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
