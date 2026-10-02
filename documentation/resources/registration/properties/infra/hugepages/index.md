---
page_title: "infra.hugepages"
subcategory: ""
description: "Hugepage settings for CE on K8s SMV2 site."
xcsh_docs: {"aliases": ["infra hugepages"], "body_bytes": 2243, "body_sha256": "sha256:12e8e53fe879dfde347d37ddaee4ffed647348c1ab898821ace90692b2aa4705", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hugepages", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "documentation/resources/registration/properties/infra/hugepages/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3201100030132010-0301200013320120-2133223300001321-2020020110003202-2111033011003210-3210020322302302-2121023211021101-0213121211230210", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hugepages"], "schema_version": 1, "sections": [{"aliases": ["free"], "anchor": "schema-infra--hugepages--free", "description": "Total number of free hugepages present.", "document_id": "xcsh-docs:resources:registration:properties:infra:hugepages", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hugepages", "free"], "syntax": "attribute", "type": "number"}, {"aliases": ["page size"], "anchor": "schema-infra--hugepages--page_size", "description": "Size of each hugepage.", "document_id": "xcsh-docs:resources:registration:properties:infra:hugepages", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hugepages", "page_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["total"], "anchor": "schema-infra--hugepages--total", "description": "Total number of hugepages present.", "document_id": "xcsh-docs:resources:registration:properties:infra:hugepages", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hugepages", "total"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hugepages/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Hugepage settings for CE on K8s SMV2 site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hugepages

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- infra.hugepages

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Hugepage settings for CE on K8s SMV2 site.

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

Terraform syntax:

```terraform
hugepages {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hugepages--free"></a>

### free property

Type: `"number"`. Optional.

Free Hugepages. Total number of free hugepages present.

Upstream description:

Total number of free hugepages present.

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

<a id="schema-infra--hugepages--page_size"></a>

### page_size property

Type: `"number"`. Optional.

Hugepage Size. Size of each hugepage.

Upstream description:

Size of each hugepage.

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

<a id="schema-infra--hugepages--total"></a>

### total property

Type: `"number"`. Optional.

Total Hugepages. Total number of hugepages present.

Upstream description:

Total number of hugepages present.

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

## Next pages

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
