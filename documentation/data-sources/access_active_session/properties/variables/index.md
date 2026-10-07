---
page_title: "variables"
subcategory: ""
description: "Variables. Session variables as key-value pairs."
xcsh_docs: {"aliases": ["variables"], "body_bytes": 792, "body_sha256": "sha256:0d4d80f605076e9846efa22360c00a46f713efc23729439bd666996290200aa6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_session:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_session:properties:variables", "parent_id": "xcsh-docs:data-sources:access_active_session:reference", "path": "documentation/data-sources/access_active_session/properties/variables/index.md", "product": "distributed-cloud", "provider_name": "access_active_session", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3332213132323331-1233202310231200-1020203011130331-3330110313102123-1131222000312201-2013021303031111-1011112331000012-2101322230230113", "registry_path": "docs/guides/data-sources--access_active_session--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["variables"], "schema_version": 1, "sections": [{"aliases": ["variables value"], "anchor": "schema-variables--value", "description": "Variable Value. The value of the session variable.", "document_id": "xcsh-docs:data-sources:access_active_session:properties:variables", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["variables", "value"], "syntax": "attribute", "type": "string"}, {"aliases": ["variables variable"], "anchor": "schema-variables--variable", "description": "Variable Name. The name of the session variable.", "document_id": "xcsh-docs:data-sources:access_active_session:properties:variables", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["variables", "variable"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_session/properties/variables/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Variables. Session variables as key-value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# variables

Breadcrumbs:

- [xcsh_access_active_session](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/)
- variables

<a id="section"></a>

Type: `"list"`. Computed.

Variables. Session variables as key-value pairs.

## Direct properties

<a id="schema-variables--value"></a>

### value property

Type: `"string"`. Computed.

Variable Value. The value of the session variable.

<a id="schema-variables--variable"></a>

### variable property

Type: `"string"`. Computed.

Variable Name. The name of the session variable.
