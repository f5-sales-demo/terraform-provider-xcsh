---
page_title: "cve_ids"
subcategory: ""
description: "cve_ids for xcsh_waf_threats."
xcsh_docs: {"aliases": [], "body_bytes": 949, "body_sha256": "sha256:90d3a683402b72e0ad788a0f58866642c5dbd39575b93ab3ba4581dbcfc83498", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:properties:cve_ids", "parent_id": "xcsh-docs:data-sources:waf_threats:reference", "path": "documentation/data-sources/waf_threats/properties/cve_ids/index.md", "provider_name": "waf_threats", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["cve_ids"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/properties/cve_ids/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cve_ids for xcsh_waf_threats.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
