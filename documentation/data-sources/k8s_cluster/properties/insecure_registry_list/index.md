---
page_title: "insecure_registry_list"
subcategory: ""
description: "List of docker insecure registries."
xcsh_docs: {"aliases": ["insecure registry list"], "body_bytes": 2765, "body_sha256": "sha256:89694ba8012b4b75db2e4a8ebd564f3f29bf0c95555d580d51b925d412b74004", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:insecure_registry_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/insecure_registry_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2000030331132132-2030211331113330-3220120113033231-2122210221011231-0213200233103132-2313012130101320-3322020233012010-0101023233133021", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["insecure_registry_list"], "schema_version": 1, "sections": [{"aliases": ["insecure registry list insecure registries"], "anchor": "schema-insecure_registry_list--insecure_registries", "description": "List of docker insecure registries in format \"example.com:5000\"", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:insecure_registry_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["insecure_registry_list", "insecure_registries"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/insecure_registry_list/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "List of docker insecure registries.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# insecure_registry_list

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- insecure_registry_list

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: insecure\_registry\_list, no\_insecure\_registries; Default: no\_insecure\_registries\]
Docker Insecure Registry List. List of docker insecure registries.

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

OneOf alternatives in this subsection:

- [insecure_registry_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/insecure_registry_list/#section)
- [no_insecure_registries](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/no_insecure_registries/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-insecure_registry_list--insecure_registries"></a>

### insecure_registries property

Type: `["list", "string"]`. Computed.

List of docker insecure registries in format 'example.com:5000'.

Additional upstream details:

List of docker insecure registries in format "example.com:5000"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
