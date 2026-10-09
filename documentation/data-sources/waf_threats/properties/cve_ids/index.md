---
page_title: "cve_ids"
subcategory: ""
description: "CVE ID List. A list of CVE IDs."
xcsh_docs: {"aliases": ["cve ids"], "body_bytes": 742, "body_sha256": "sha256:37151db5542f5e39cd548563ae3926cfa13bd31e34d004b32ce5b88b0e8fcbae", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "parent_id": "xcsh-docs:data-sources:waf_threats:reference", "path": "documentation/data-sources/waf_threats/properties/cve_ids/index.md", "product": "distributed-cloud", "provider_name": "waf_threats", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3310211112211330-1200201202322030-3033320213030222-0131133002322122-0001312021112313-3301333212123330-1201120120331021-1220322301000111", "registry_path": "docs/guides/data-sources--waf_threats--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cve_ids"], "schema_version": 1, "sections": [{"aliases": ["cve ids ids"], "anchor": "schema-cve_ids--ids", "description": "IDs. A list of CVE IDs.", "document_id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cve_ids", "ids"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/properties/cve_ids/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "CVE ID List. A list of CVE IDs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```
