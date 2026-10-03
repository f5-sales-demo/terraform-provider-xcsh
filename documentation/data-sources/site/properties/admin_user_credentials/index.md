---
page_title: "admin_user_credentials"
subcategory: "Infrastructure"
description: "Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin' user will be setup and customers can access these nodes via either the node local WebUI or via SSH to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access."
xcsh_docs: {"aliases": ["admin user credentials", "authentication", "credential setup", "credentials"], "body_bytes": 1557, "body_sha256": "sha256:955b9e853622e15b666e0fad27b7ace087ee8ee91a8c92e2105a074ddf9a0e92", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:admin_user_credentials", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/admin_user_credentials/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0203302231220301-0223312203020023-1332333012202212-0113313102210020-3313032202232002-1201330310223133-3010212020223202-0231210121213032", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["admin_user_credentials"], "schema_version": 1, "sections": [{"aliases": ["admin user credentials admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["admin_user_credentials", "admin_password"], "syntax": "attribute", "type": "object"}, {"aliases": ["admin user credentials ssh key"], "anchor": "schema-admin_user_credentials--ssh_key", "description": "Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can SSH to the nodes of this Customer Edge site using admin as the user.", "document_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["admin_user_credentials", "ssh_key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/admin_user_credentials/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin' user will be setup and customers can access these nodes via either the node local WebUI or via SSH to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_user_credentials

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- admin_user_credentials

<a id="section"></a>

Type: `"single"`. Computed.

Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin'
user will be setup and customers can access these nodes via either the node local WebUI or via SSH
to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/admin_user_credentials/admin_password/): complete subsection reference.

<a id="schema-admin_user_credentials--ssh_key"></a>

### ssh_key property

Type: `"string"`. Computed.

Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can
SSH to the nodes of this Customer Edge site using admin as the user.

## Next pages

- [admin_user_credentials.admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/admin_user_credentials/admin_password/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
