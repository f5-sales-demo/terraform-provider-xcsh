---
page_title: "dataplane_ref.nginx_instance"
subcategory: ""
description: "dataplane_ref.nginx_instance for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 1502, "body_sha256": "sha256:eec7e5aaaba2d986246081e097d0ac67ce07aa2b0b067ef81ae53df9f70e60a3", "canonical_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref:nginx_instance", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:dataplane_ref", "path": "docs/guides/data-sources--nginx_server--properties--dataplane_ref--nginx_instance.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dataplane_ref", "nginx_instance"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/dataplane_ref/nginx_instance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dataplane_ref.nginx_instance for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dataplane_ref.nginx_instance

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- [Property reference](data-sources--nginx_server--reference.md)
- [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md)
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

- [dataplane_ref](data-sources--nginx_server--properties--dataplane_ref.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
