---
page_title: "infra.hw_info.kernel"
subcategory: ""
description: "Kernel information."
xcsh_docs: {"aliases": ["infra hw info kernel"], "body_bytes": 2972, "body_sha256": "sha256:a4e2aa2f8856101d6adaeed619d097a99dc18643b25290252f2229617041f926", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "documentation/data-sources/registration/properties/infra/hw_info/kernel/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3221100101322111-0322230203320111-2222313030300303-2111330331331222-3133232133112212-1023103301201132-2202120013320222-3133222303011322", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "kernel"], "schema_version": 1, "sections": [{"aliases": ["architecture"], "anchor": "schema-infra--hw_info--kernel--architecture", "description": "Kernel architecture.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "kernel", "architecture"], "syntax": "attribute", "type": "string"}, {"aliases": ["release"], "anchor": "schema-infra--hw_info--kernel--release", "description": "Kernel release.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "kernel", "release"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-infra--hw_info--kernel--version", "description": "Kernel version.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "kernel", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/kernel/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Kernel information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.kernel

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- infra.hw_info.kernel

<a id="section"></a>

Type: `"single"`. Computed.

Kernel. Kernel information.

Upstream description:

Kernel information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-infra--hw_info--kernel--architecture"></a>

### architecture property

Type: `"string"`. Computed.

Architecture. Kernel architecture.

Upstream description:

Kernel architecture.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--kernel--release"></a>

### release property

Type: `"string"`. Computed.

Release. Kernel release.

Upstream description:

Kernel release.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--kernel--version"></a>

### version property

Type: `"string"`. Computed.

Version. Kernel version.

Upstream description:

Kernel version.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
