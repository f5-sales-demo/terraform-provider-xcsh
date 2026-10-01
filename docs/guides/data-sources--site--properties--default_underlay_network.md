---
page_title: "default_underlay_network"
subcategory: "Infrastructure"
description: "default_underlay_network for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1144, "body_sha256": "sha256:50bcbf82932f9015bfdfce10d1c5cad2bd28b4f701236b43a0d0398be4a3e9ef", "canonical_id": "xcsh-docs:data-sources:site:properties:default_underlay_network", "child_ids": ["xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside"], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:default_underlay_network", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "docs/guides/data-sources--site--properties--default_underlay_network.md", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_underlay_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/default_underlay_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_underlay_network for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_underlay_network

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- default_underlay_network

<a id="section"></a>

Type: `"single"`. Computed.

Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP
tunnels for DC Cluster Group) Default is site-local-outside network.

## Direct properties

- [site_local_inside](data-sources--site--properties--default_underlay_network--site_local_inside.md): complete subsection reference.

- [site_local_outside](data-sources--site--properties--default_underlay_network--site_local_outside.md): complete subsection reference.

## Next pages

- [default_underlay_network.site_local_inside](data-sources--site--properties--default_underlay_network--site_local_inside.md)
- [default_underlay_network.site_local_outside](data-sources--site--properties--default_underlay_network--site_local_outside.md)
- [Property reference](data-sources--site--reference.md)
- [xcsh_site](../data-sources/site.md)
