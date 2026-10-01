---
page_title: "re_select.specific_re"
subcategory: "Infrastructure"
description: "re_select.specific_re for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 990, "body_sha256": "sha256:9b5196bc360c8a5ea729bbcebeebbd7d401c18a8be09cf071003c992e0296ee2", "canonical_id": "xcsh-docs:data-sources:site:properties:re_select:specific_re", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:re_select:specific_re", "parent_id": "xcsh-docs:data-sources:site:properties:re_select", "path": "docs/guides/data-sources--site--properties--re_select--specific_re.md", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_select", "specific_re"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/re_select/specific_re/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_select.specific_re for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select.specific_re

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- [re_select](data-sources--site--properties--re_select.md)
- re_select.specific_re

<a id="section"></a>

Type: `"single"`. Computed.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

## Direct properties

<a id="schema-re_select--specific_re--backup_re"></a>

### backup_re property

Type: `"string"`. Computed.

Select backup RE for this site, cannot be the same as Primary RE.

<a id="schema-re_select--specific_re--primary_re"></a>

### primary_re property

Type: `"string"`. Computed.

Primary RE Geography. Select primary RE for this site.

## Next pages

- [re_select](data-sources--site--properties--re_select.md)
- [xcsh_site](../data-sources/site.md)
