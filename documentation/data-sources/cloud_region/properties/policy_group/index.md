---
page_title: "policy_group"
subcategory: ""
description: "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["policy group"], "body_bytes": 1540, "body_sha256": "sha256:0d63402a5c27c826985854a6495b038aea32706cd86dd3d3c94736b68d720462", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "parent_id": "xcsh-docs:data-sources:cloud_region:reference", "path": "documentation/data-sources/cloud_region/properties/policy_group/index.md", "product": "distributed-cloud", "provider_name": "cloud_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3101233321321233-2033200000010013-1303333303233320-1302030121330213-1030221132230033-0223111023331212-0322020312131322-0312230310312332", "registry_path": "docs/guides/data-sources--cloud_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_group"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-policy_group--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_group", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-policy_group--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_group", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-policy_group--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_group", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/properties/policy_group/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_group

Breadcrumbs:

- [xcsh_cloud_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/)
- policy_group

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-policy_group--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-policy_group--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-policy_group--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/)
- [xcsh_cloud_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/)
