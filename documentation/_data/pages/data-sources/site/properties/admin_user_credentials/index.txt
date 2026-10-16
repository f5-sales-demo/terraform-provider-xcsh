---
page_title: "admin_user_credentials"
subcategory: "Infrastructure"
description: "Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin' user will be setup and customers can access these nodes via either the node local WebUI or via SSH to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access."
xcsh_docs: {"aliases": ["admin user credentials", "authentication", "credential setup", "credentials"], "body_bytes": 1174, "body_sha256": "sha256:e33ffaccdc39c755c67179b882dbc8268a2a21fd7c0c8d4393e8666f89147876", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:admin_user_credentials", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/admin_user_credentials/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0203302231220301-0223312203020023-1332333012202212-0113313102210020-3313032202232002-1201330310223133-3010212020223202-0231210121213032", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["admin_user_credentials"], "schema_version": 1, "sections": [{"aliases": ["admin user credentials admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["admin_user_credentials", "admin_password"], "syntax": "attribute", "type": "object"}, {"aliases": ["admin user credentials ssh key"], "anchor": "schema-admin_user_credentials--ssh_key", "description": "Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can SSH to the nodes of this Customer Edge site using admin as the user.", "document_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["admin_user_credentials", "ssh_key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/admin_user_credentials/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin' user will be setup and customers can access these nodes via either the node local WebUI or via SSH to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
