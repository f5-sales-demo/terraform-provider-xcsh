---
page_title: "dataplane_ref.nginx_csg"
subcategory: ""
description: "dataplane_ref.nginx_csg for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 1477, "body_sha256": "sha256:e2cbdf8a33f543a26b2933789b14fdd05a19a061b2c41afd20b0fabaa2e3847e", "canonical_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_csg", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_csg", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "path": "docs/guides/data-sources--nginx_server--properties--dataplane_ref--nginx_csg.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dataplane_ref", "nginx_csg"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/dataplane_ref/nginx_csg/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dataplane_ref.nginx_csg for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dataplane_ref.nginx_csg

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- [Property reference](data-sources--nginx_server--reference.md)
- [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md)
- dataplane_ref.nginx_csg

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-dataplane_ref--nginx_csg--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-dataplane_ref--nginx_csg--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-dataplane_ref--nginx_csg--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

## Next pages

- [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
