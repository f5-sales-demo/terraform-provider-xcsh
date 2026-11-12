---
page_title: "dataplane_ref"
subcategory: ""
description: "DataplaneReference."
xcsh_docs: {"aliases": ["dataplane ref"], "body_bytes": 806, "body_sha256": "sha256:85e89fe21f4f4d95103428594066e4565ce17ce1043b1a352b57d368a403b805", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_csg", "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "parent_id": "xcsh-docs:data-sources:nginx_server:reference", "path": "documentation/data-sources/nginx_server/properties/dataplane_ref/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-0220302120020200-1121031133010013-1221013203013032-3302322122311212-3303003030302120-3213212133131003-1233321021300112-3333301303233231", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dataplane_ref"], "schema_version": 1, "sections": [{"aliases": ["dataplane ref nginx csg"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_csg", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dataplane_ref", "nginx_csg"], "syntax": "attribute", "type": "object"}, {"aliases": ["dataplane ref nginx instance"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dataplane_ref", "nginx_instance"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/dataplane_ref/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "DataplaneReference.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dataplane_ref

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- dataplane_ref

<a id="section"></a>

Type: `"single"`. Computed.

DataplaneReference.

## Direct properties

- [nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_csg/): complete subsection reference.

- [nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/): complete subsection reference.
