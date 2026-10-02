---
page_title: "cve_ids"
subcategory: ""
description: "CVE ID List. A list of CVE IDs."
xcsh_docs: {"aliases": ["cve ids"], "body_bytes": 949, "body_sha256": "sha256:90d3a683402b72e0ad788a0f58866642c5dbd39575b93ab3ba4581dbcfc83498", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "parent_id": "xcsh-docs:data-sources:waf_threats:reference", "path": "documentation/data-sources/waf_threats/properties/cve_ids/index.md", "product": "distributed-cloud", "provider_name": "waf_threats", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3310211112211330-1200201202322030-3033320213030222-0131133002322122-0001312021112313-3301333212123330-1201120120331021-1220322301000111", "registry_path": "docs/guides/data-sources--waf_threats--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cve_ids"], "schema_version": 1, "sections": [{"aliases": ["ids"], "anchor": "schema-cve_ids--ids", "description": "IDs. A list of CVE IDs.", "document_id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cve_ids", "ids"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/properties/cve_ids/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "CVE ID List. A list of CVE IDs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cve_ids

Breadcrumbs:

- [xcsh_waf_threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/)
- cve_ids

<a id="section"></a>

Type: `"single"`. Optional.

CVE ID List. A list of CVE IDs.

## Direct properties

<a id="schema-cve_ids--ids"></a>

### ids property

Type: `["list", "string"]`. Optional.

IDs. A list of CVE IDs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/properties/)
- [xcsh_waf_threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/)
