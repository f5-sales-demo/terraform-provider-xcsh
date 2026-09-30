---
page_title: "main_nodes"
subcategory: "Infrastructure"
description: "main_nodes for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 922, "body_sha256": "sha256:dd6c70f3bdaa6d140771fc566505e2f1bc7adda0051d76a199d02a4cd7ef65a8", "canonical_id": "xcsh-docs:data-sources:site:properties:main_nodes", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:main_nodes", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "docs/guides/data-sources--site--properties--main_nodes.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["main_nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/main_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "main_nodes for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# main_nodes

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- main_nodes

<a id="section"></a>

Type: `"list"`. Computed.

Connectivity information of main/master nodes to create a full mesh of Phobos services across all
CEs in a site-mesh-group or dc-cluster-group.

## Direct properties

<a id="schema-main_nodes--name"></a>

### name property

Type: `"string"`. Computed.

Name of the master/main node on the site.

<a id="schema-main_nodes--sli_address"></a>

### sli_address property

Type: `"string"`. Computed.

Site Local Inside IP addresses. Site Local Inside IP address.

<a id="schema-main_nodes--slo_address"></a>

### slo_address property

Type: `"string"`. Computed.

Site Local Outside IP addresses. Site Local Outside IP address.

## Next pages

- [Property reference](data-sources--site--reference.md)
- [xcsh_site](../data-sources/site.md)
