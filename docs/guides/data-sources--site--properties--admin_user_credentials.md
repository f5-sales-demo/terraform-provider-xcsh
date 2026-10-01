---
page_title: "admin_user_credentials"
subcategory: "Infrastructure"
description: "admin_user_credentials for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1249, "body_sha256": "sha256:64ef60af35ed52dce8b4f0a7380ddcb6107cab082a75199b296203f2796224d4", "canonical_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials", "child_ids": ["xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password"], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:admin_user_credentials", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "docs/guides/data-sources--site--properties--admin_user_credentials.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["admin_user_credentials"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/admin_user_credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "admin_user_credentials for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_user_credentials

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- admin_user_credentials

<a id="section"></a>

Type: `"single"`. Computed.

Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin'
user will be setup and customers can access these nodes via either the node local WebUI or via SSH
to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.

## Direct properties

- [admin_password](data-sources--site--properties--admin_user_credentials--admin_password.md): complete subsection reference.

<a id="schema-admin_user_credentials--ssh_key"></a>

### ssh_key property

Type: `"string"`. Computed.

Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can
SSH to the nodes of this Customer Edge site using admin as the user.

## Next pages

- [admin_user_credentials.admin_password](data-sources--site--properties--admin_user_credentials--admin_password.md)
- [Property reference](data-sources--site--reference.md)
- [xcsh_site](../data-sources/site.md)
