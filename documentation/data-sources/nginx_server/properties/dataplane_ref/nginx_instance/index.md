---
page_title: "dataplane_ref.nginx_instance"
subcategory: ""
description: "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["dataplane ref nginx instance"], "body_bytes": 1759, "body_sha256": "sha256:d45f112179811d80e0d4863b803eaabed68fbb3f83fe0350be1c4e54df40370a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "path": "documentation/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0212310303131230-0100121111300131-1310001000023210-3111302110003122-2000320203202330-3032212130011233-2001300010333233-2322333200110113", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dataplane_ref", "nginx_instance"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-dataplane_ref--nginx_instance--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dataplane_ref", "nginx_instance", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-dataplane_ref--nginx_instance--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dataplane_ref", "nginx_instance", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-dataplane_ref--nginx_instance--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dataplane_ref", "nginx_instance", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dataplane_ref.nginx_instance

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- [dataplane_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/)
- dataplane_ref.nginx_instance

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-dataplane_ref--nginx_instance--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-dataplane_ref--nginx_instance--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-dataplane_ref--nginx_instance--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

## Next pages

- [dataplane_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/dataplane_ref/)
- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
