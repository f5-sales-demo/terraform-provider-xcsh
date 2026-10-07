---
page_title: "nodes"
subcategory: ""
description: "nodes for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": ["nodes"], "body_bytes": 1389, "body_sha256": "sha256:262226a75ce4e614b40a417a89e04dca4a887349606c8e2b8fcb46e0cd7a27fc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "path": "documentation/data-sources/smsv2_aws_runtime/properties/nodes/index.md", "product": "distributed-cloud", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3100030321300033-0301332220323222-3011133301220213-2132001321111023-1131210221212320-3011001021211332-0322233030121220-1001203131220132", "registry_path": "docs/guides/data-sources--smsv2_aws_runtime--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["nodes"], "schema_version": 1, "sections": [{"aliases": ["nodes mac"], "anchor": "schema-nodes--mac", "description": "mac", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "enum_extraction_complete": true, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes", "mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["nodes node"], "anchor": "schema-nodes--node", "description": "node", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "enum_extraction_complete": true, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["nodes role"], "anchor": "schema-nodes--role", "description": "role", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["sli", "slo"], "version": 1}], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes", "role"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/properties/nodes/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "nodes for xcsh_smsv2_aws_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nodes

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/)
- nodes

<a id="section"></a>

Type: `"map"`. Required.

## Direct properties

<a id="schema-nodes--mac"></a>

### mac property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-nodes--node"></a>

### node property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-nodes--role"></a>

### role property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["sli","slo"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```
