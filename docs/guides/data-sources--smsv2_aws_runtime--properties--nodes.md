---
page_title: "nodes"
subcategory: ""
description: "nodes for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 1031, "body_sha256": "sha256:9cf8aa40a76de92d34c9157e573e4b679f32db8cb61e68ddd89af941381c06f3", "canonical_id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "child_ids": [], "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:properties:nodes", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:reference", "path": "docs/guides/data-sources--smsv2_aws_runtime--properties--nodes.md", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/properties/nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nodes for xcsh_smsv2_aws_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# nodes

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md)
- [Property reference](data-sources--smsv2_aws_runtime--reference.md)
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

- [Property reference](data-sources--smsv2_aws_runtime--reference.md)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md)
