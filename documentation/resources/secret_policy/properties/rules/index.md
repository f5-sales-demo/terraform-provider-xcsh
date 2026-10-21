---
page_title: "rules"
subcategory: "Security"
description: "rules for xcsh_secret_policy."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 1765, "body_sha256": "sha256:761d20b6c1e1462d4c0838a7cb43ab32c2a0dab079a875f0243c20398e34cd0d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:secret_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_policy:properties:rules", "parent_id": "xcsh-docs:resources:secret_policy:reference", "path": "documentation/resources/secret_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "secret_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1111232303130011-1030301132230333-3012212320310123-1213022030003030-3202032333121311-2232110321210330-3200001233322211-0333333311023332", "registry_path": "docs/guides/resources--secret_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["rules action"], "anchor": "schema-rules--action", "description": "action", "document_id": "xcsh-docs:resources:secret_policy:properties:rules", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ALLOW", "DENY"], "version": 1}], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules client selector"], "anchor": "schema-rules--client_selector", "description": "Kubernetes selector expression. Omitted selector matches every client; use only for a final DENY rule.", "document_id": "xcsh-docs:resources:secret_policy:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_selector"], "syntax": "attribute", "type": "list"}, {"aliases": ["rules description"], "anchor": "schema-rules--description", "description": "description", "document_id": "xcsh-docs:resources:secret_policy:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "description"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules name"], "anchor": "schema-rules--name", "description": "name", "document_id": "xcsh-docs:resources:secret_policy:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "rules for xcsh_secret_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["secret_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_secret_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/)
- rules

<a id="section"></a>

Type: `"list"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{listvalidator.SizeBetween(1, 256)}
```

## Direct properties

<a id="schema-rules--action"></a>

### action property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("ALLOW",
    "DENY")}
```

<a id="schema-rules--client_selector"></a>

### client_selector property

Type: `["list", "string"]`. Optional.

Kubernetes selector expression. Omitted selector matches every client; use only for a final DENY
rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{listvalidator.SizeAtMost(1)}
```

<a id="schema-rules--description"></a>

### description property

Type: `"string"`. Optional.

<a id="schema-rules--name"></a>

### name property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{validators.NameValidator()}
```
