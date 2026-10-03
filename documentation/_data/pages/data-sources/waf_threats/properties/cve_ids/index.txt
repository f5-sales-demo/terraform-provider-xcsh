---
page_title: "cve_ids"
subcategory: ""
description: "CVE ID List. A list of CVE IDs."
xcsh_docs: {"aliases": ["cve ids"], "body_bytes": 949, "body_sha256": "sha256:90d3a683402b72e0ad788a0f58866642c5dbd39575b93ab3ba4581dbcfc83498", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "parent_id": "xcsh-docs:data-sources:waf_threats:reference", "path": "documentation/data-sources/waf_threats/properties/cve_ids/index.md", "product": "distributed-cloud", "provider_name": "waf_threats", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3310211112211330-1200201202322030-3033320213030222-0131133002322122-0001312021112313-3301333212123330-1201120120331021-1220322301000111", "registry_path": "docs/guides/data-sources--waf_threats--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cve_ids"], "schema_version": 1, "sections": [{"aliases": ["cve ids ids"], "anchor": "schema-cve_ids--ids", "description": "IDs. A list of CVE IDs.", "document_id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cve_ids", "ids"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/properties/cve_ids/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "CVE ID List. A list of CVE IDs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": [], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
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
