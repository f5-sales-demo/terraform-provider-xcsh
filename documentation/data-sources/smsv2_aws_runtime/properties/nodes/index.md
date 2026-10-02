---
page_title: "nodes"
subcategory: ""
description: "nodes for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": ["nodes"], "body_bytes": 1338, "body_sha256": "sha256:f8bf39af3100c1d6d0962581cbf731dfd9bcea53f359d6bc8cf9dcf956ceea34", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "path": "documentation/data-sources/smsv2_aws_runtime/properties/nodes/index.md", "product": "distributed-cloud", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3100030321300033-0301332220323222-3011133301220213-2132001321111023-1131210221212320-3011001021211332-0322233030121220-1001203131220132", "registry_path": "docs/guides/data-sources--smsv2_aws_runtime--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["nodes"], "schema_version": 1, "sections": [{"aliases": ["mac"], "anchor": "schema-nodes--mac", "description": "mac", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes", "mac"], "syntax": "attribute", "type": "string"}, {"aliases": ["node"], "anchor": "schema-nodes--node", "description": "node", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["role"], "anchor": "schema-nodes--role", "description": "role", "document_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes", "role"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/properties/nodes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "nodes for xcsh_smsv2_aws_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-nodes--node"></a>

### node property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-nodes--role"></a>

### role property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/)
- [xcsh_smsv2_aws_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/)
